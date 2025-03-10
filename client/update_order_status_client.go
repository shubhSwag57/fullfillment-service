package main

import (
	"DeliveryService/pb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewDeliveryServiceClient(conn)

	// Update the status of order ID 1
	updateOrderStatus(client, 12, "DELIVERED")
}

func updateOrderStatus(client pb.DeliveryServiceClient, orderID int64, status string) {
	req := &pb.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  status,
	}

	resp, err := client.UpdateOrderStatus(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to update order status: %v", err)
	}

	log.Printf("Order status updated successfully: %s", resp.Message)
}
