package service

import (
	"DeliveryService/db"
	"DeliveryService/models"
	"DeliveryService/pb"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

// DeliveryServiceImpl implements pb.DeliveryServiceServer
type DeliveryServiceImpl struct {
	pb.UnimplementedDeliveryServiceServer
}

// AssignOrder finds an available delivery person and assigns an order
func (s *DeliveryServiceImpl) AssignOrder(ctx context.Context, req *pb.AssignOrderRequest) (*pb.AssignOrderResponse, error) {
	var deliveryPerson models.DeliveryPerson
	result := db.DB.Where("available = ?", true).First(&deliveryPerson)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("no available delivery person found")
	}

	order := models.Order{
		ID:               req.OrderId,
		DeliveryPersonID: deliveryPerson.ID,
		Status:           "ASSIGNED",
	}
	db.DB.Create(&order)

	// Mark the delivery person as unavailable
	db.DB.Model(&deliveryPerson).Update("available", false)

	return &pb.AssignOrderResponse{
		Message:          "Order assigned successfully",
		DeliveryPersonId: deliveryPerson.ID,
	}, nil
}

// UpdateOrderStatus updates an order's status
func (s *DeliveryServiceImpl) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	result := db.DB.Model(&models.Order{}).Where("id = ?", req.OrderId).Update("status", req.Status)

	if result.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	// If order is completed, mark delivery person as available
	orderID := req.GetOrderId()
	status := req.GetStatus()

	url := fmt.Sprintf("http://localhost:8082/api/orders/%d/status?status=%s", orderID, status)
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	//client := &http.Client{}
	//
	//// ✅ Create the HTTP request
	//reqBody := strings.NewReader("") // No body needed for PUT request
	//request, err := http.NewRequest("PUT", url, reqBody)
	//if err != nil {
	//	log.Printf("Error creating request: %v", err)
	//	return nil, err
	//}

	// ✅ Set Basic Authentication header (This is the Fix)
	username := "admin"
	password := "admin123"
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	request.Header.Set("Authorization", "Basic "+auth)

	resp, err := client.Do(request)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return nil, err
	}

	log.Printf("Order update response: %s", string(body))

	if req.Status == "DELIVERED" {
		var order models.Order
		db.DB.First(&order, req.OrderId)
		db.DB.Model(&models.DeliveryPerson{}).Where("id = ?", order.DeliveryPersonID).Update("available", true)
	}

	return &pb.UpdateOrderStatusResponse{
		Message: "Order status updated successfully",
	}, nil
}

func (s *DeliveryServiceImpl) RegisterDeliveryPerson(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	deliveryPerson := models.DeliveryPerson{
		Name:      req.Name,
		Password:  string(hashedPassword),
		Available: true,
	}
	if err := db.DB.Create(&deliveryPerson).Error; err != nil {
		return nil, errors.New("name already taken")
	}

	return &pb.RegisterResponse{Message: "Registration successful"}, nil
}

func (s *DeliveryServiceImpl) LoginDeliveryPerson(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var deliveryPerson models.DeliveryPerson
	if err := db.DB.Where("name = ?", req.Name).First(&deliveryPerson).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(deliveryPerson.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &pb.LoginResponse{Id: deliveryPerson.ID, Message: "Login successful"}, nil
}

func (s *DeliveryServiceImpl) UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.UpdateLocationResponse, error) {
	var deliveryPerson models.DeliveryPerson
	if err := db.DB.First(&deliveryPerson, req.Id).Error; err != nil {
		return nil, errors.New("delivery person not found")
	}

	db.DB.Model(&deliveryPerson).Updates(models.DeliveryPerson{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	})

	return &pb.UpdateLocationResponse{Message: "Location updated successfully"}, nil
}
