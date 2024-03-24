// gateway-service/cmd/gateway/main.go

package main

import (
	"log"
	"net/http"

	"gateway/internal/gateway"
	"gateway/pkg/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("hashing:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to hashing service: %v", err)
	}
	defer conn.Close()

	hashingClient := pb.NewHashingServiceClient(conn)

	gatewayService := gateway.NewService(hashingClient)

	http.HandleFunc("/checkhash", gatewayService.CheckHashHandler)
	http.HandleFunc("/gethash", gatewayService.GetHashHandler)
	http.HandleFunc("/createhash", gatewayService.CreateHashHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
