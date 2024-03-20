package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"hashing/internal/hashing"
	"hashing/pkg/pb"
)

func main() {
	server := grpc.NewServer()

	pb.RegisterHashingServiceServer(server, hashing.NewHashingService())

	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
