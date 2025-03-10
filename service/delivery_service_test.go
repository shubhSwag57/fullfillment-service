package service_test

import (
	"DeliveryService/db"
	"DeliveryService/models"
	"DeliveryService/pb"
	"DeliveryService/service"
	"context"
	_ "errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func setupTestDB() {
	db.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.DB.AutoMigrate(&models.DeliveryPerson{}, &models.Order{})
}

func TestRegisterDeliveryPerson(t *testing.T) {
	setupTestDB()
	service := &service.DeliveryServiceImpl{}

	req := &pb.RegisterRequest{
		Name:     "John",
		Password: "password123",
	}
	resp, err := service.RegisterDeliveryPerson(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Registration successful", resp.Message)
}

func TestLoginDeliveryPerson(t *testing.T) {
	setupTestDB()
	service := &service.DeliveryServiceImpl{}

	deliveryPerson := models.DeliveryPerson{
		Name:     "John",
		Password: "$2a$10$7QFEC2EIxRSOpK.PM6YSe.SBCOPNOH6xTqQ5wY/fO/lo7J2J1S2Ni", // bcrypt hash for 'password123'
	}
	db.DB.Create(&deliveryPerson)

	req := &pb.LoginRequest{
		Name:     "John",
		Password: "password123",
	}
	resp, err := service.LoginDeliveryPerson(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, deliveryPerson.ID, resp.Id)
}

func TestAssignOrder(t *testing.T) {
	setupTestDB()
	service := &service.DeliveryServiceImpl{}

	deliveryPerson := models.DeliveryPerson{
		Name:      "John",
		Password:  "password",
		Available: true,
	}
	db.DB.Create(&deliveryPerson)

	req := &pb.AssignOrderRequest{
		OrderId: 1,
	}
	resp, err := service.AssignOrder(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, deliveryPerson.ID, resp.DeliveryPersonId)
}

func TestUpdateOrderStatus(t *testing.T) {
	setupTestDB()
	service := &service.DeliveryServiceImpl{}

	order := models.Order{
		ID:               1,
		DeliveryPersonID: 1,
		Status:           "ASSIGNED",
	}
	db.DB.Create(&order)

	req := &pb.UpdateOrderStatusRequest{
		OrderId: 1,
		Status:  "DELIVERED",
	}
	resp, err := service.UpdateOrderStatus(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Order status updated successfully", resp.Message)
}

func TestUpdateLocation(t *testing.T) {
	setupTestDB()
	service := &service.DeliveryServiceImpl{}

	deliveryPerson := models.DeliveryPerson{
		Name:      "John",
		Password:  "password",
		Latitude:  0.0,
		Longitude: 0.0,
	}
	db.DB.Create(&deliveryPerson)

	req := &pb.UpdateLocationRequest{
		Id:        deliveryPerson.ID,
		Latitude:  10.0,
		Longitude: 20.0,
	}
	resp, err := service.UpdateLocation(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Location updated successfully", resp.Message)
}
