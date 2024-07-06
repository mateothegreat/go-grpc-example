package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	_, err := NewServer(ctx, "0.0.0.0:18880")
	if err != nil {
		t.Errorf("failed to create server: %v", err)
	}

	// Capture interrupt signals to cancel the context and gracefully shut down the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Received shutdown signal")
		cancel()
	}()

	// Run the server indefinitely
	<-ctx.Done()

	// Give the server some time to shutdown gracefully
	time.Sleep(2 * time.Second)
	fmt.Println("Server shutdown complete")
}
