package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bsarvan/productInfo/service/productInfopb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	fmt.Println("Starting the Server\n")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	s := grpc.NewServer()
	productInfopb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port: %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error to serve: %v", err)
	}
}
