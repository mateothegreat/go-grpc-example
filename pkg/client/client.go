package client

import (
	"fmt"
	"log"

	"github.com/cmdbstream/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(server string) (protos.ServerServiceClient, error) {
	log.Printf("Connecting to server at %s", server)
	conn, err := grpc.NewClient(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server at %s: %v", server, err)
	}

	return protos.NewServerServiceClient(conn), nil
}
