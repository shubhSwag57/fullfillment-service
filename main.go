package main

import (
	"DeliveryService/db"
	"DeliveryService/models"
	"DeliveryService/server"
	"log"
)

func main() {
	db.InitDB()
	db.RunMigrations()

	deliveryPerson := models.DeliveryPerson{
		Name:      "John Doe",
		Password:  "securepassword",
		Latitude:  37.7749,
		Longitude: -122.4194,
		Available: true,
	}

	// Insert the delivery partner into the database
	if err := db.DB.Create(&deliveryPerson).Error; err != nil {
		log.Fatalf("Failed to insert delivery partner: %v", err)
	}

	log.Println("Delivery partner inserted successfully")

	deliveryPerson = models.DeliveryPerson{
		Name:      "John Doe 1",
		Password:  "securepassword",
		Latitude:  37.774989,
		Longitude: -122.419434,
		Available: true,
	}

	// Insert the delivery partner into the database
	if err := db.DB.Create(&deliveryPerson).Error; err != nil {
		log.Fatalf("Failed to insert delivery partner: %v", err)
	}

	log.Println("Delivery partner inserted successfully")
	server.StartGRPCServer()

	// Create a new delivery partner

}
