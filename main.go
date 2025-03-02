package main

import (
	"DeliveryService/db"
	"DeliveryService/server"
)

func main() {
	db.InitDB()
	db.RunMigrations()
	server.StartGRPCServer()
}
