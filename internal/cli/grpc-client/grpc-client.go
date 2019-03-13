package client

import (
	"log"

	pb "github.com/mjlaufer/go-grpc-service/api/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

// RunClient connects a gRPC client to the server.
func RunClient() (c pb.UserServiceClient, conn *grpc.ClientConn) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	c = pb.NewUserServiceClient(conn)

	log.Printf("gRPC client connected at %v\n\n", address)

	return c, conn
}
