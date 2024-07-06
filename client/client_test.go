package client

import (
	"context"
	"log"
	"testing"

	"github.com/mateothegreat/go-grpc-example/protos"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("localhost:8888")
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}

	res, err := client.SayHello(context.Background(), &protos.ServerRequest{
		Data: "Hello, Server!",
	}, nil)
	if err != nil {
		t.Errorf("failed to call SayHello: %v", err)
	}

	log.Printf("Response: %s", res.Message)
}
