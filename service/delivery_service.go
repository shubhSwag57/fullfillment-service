package service

import (
	"DeliveryService/db"
	"DeliveryService/models"
	"DeliveryService/pb"
	"context"
	"errors"
	"gorm.io/gorm"
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
	if req.Status == "DELIVERED" {
		var order models.Order
		db.DB.First(&order, req.OrderId)
		db.DB.Model(&models.DeliveryPerson{}).Where("id = ?", order.DeliveryPersonID).Update("available", true)
	}

	return &pb.UpdateOrderStatusResponse{
		Message: "Order status updated successfully",
	}, nil
}
