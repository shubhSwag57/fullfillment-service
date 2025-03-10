package service

import (
	"DeliveryService/db"
	"DeliveryService/models"
	"DeliveryService/pb"
	"context"
	"errors"
	"fmt"
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

	url := fmt.Sprintf("http://localhost:8085/orders/%s/status?status=%s", orderID, status)
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

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
