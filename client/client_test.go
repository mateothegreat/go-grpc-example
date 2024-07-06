package client

import (
	"context"
	"log"
	"testing"

	"github.com/mateothegreat/go-grpc-example/protos"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("127.0.0.1:18880")
	if err != nil {
		t.Errorf("failed to create client: %v", err)
		return
	}

	res, err := client.SayHello(context.Background(), &protos.ServerRequest{
		Data: "Hello, Server!",
	})
	if err != nil {
		t.Errorf("failed to call SayHello: %v", err)
		return
	}

	log.Printf("Response: %s", res.Message)
}
