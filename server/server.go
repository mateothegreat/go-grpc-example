package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/mateothegreat/go-grpc-example/protos"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedServerServiceServer
	grpcServer *grpc.Server
}

func (s *Server) SayHello(ctx context.Context, in *pb.ServerRequest) (*pb.ServerResponse, error) {
	log.Printf("Received request: %v", in)
	return &pb.ServerResponse{Message: "Hello, World! "}, nil
}

func NewServer(ctx context.Context, address string) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &Server{grpcServer: grpcServer}
	pb.RegisterServerServiceServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v", err)
		}
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	// Block forever, or do something else..
	select {}
}
