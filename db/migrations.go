package db

import "DeliveryService/models"

func RunMigrations() {
	DB.AutoMigrate(&models.DeliveryPerson{}, &models.Order{})
}
