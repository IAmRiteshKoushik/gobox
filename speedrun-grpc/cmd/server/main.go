package main

import (
	"context"
	"log"
	"net"

	greetingv1 "github.com/IAmRiteshKoushik/speedrun-grpc/gen/greeting/v1"
	"google.golang.org/grpc"
)

const listenAddr = "localhost:50051"

type greetingServer struct {
	greetingv1.UnimplementedGreeterServiceServer
}

func (greetingServer) Greet(_ context.Context, request *greetingv1.GreetRequest) (*greetingv1.GreetResponse, error) {
	return &greetingv1.GreetResponse{
		Message: "Hello " + request.GetName(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("listen on %s: %v", listenAddr, err)
	}

	server := grpc.NewServer()
	// greetingv1.RegisterGreeterServiceServer(server, greetingServer{})

	log.Printf("gRPC server listening on %s", listenAddr)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("server gRPC: %v", err)
	}
}
