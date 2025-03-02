package server

import (
	"DeliveryService/pb"
	"DeliveryService/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDeliveryServiceServer(grpcServer, &service.DeliveryServiceImpl{})

	log.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
