package main

import (
	"log"
	"net"

	pb "github.com/mjlaufer/go-grpc-service/api/proto"
	"github.com/mjlaufer/go-grpc-service/internal/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register service with the gRPC server
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &app.Service{})

	log.Printf("gRPC server listening on localhost%v\n\n", port)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
