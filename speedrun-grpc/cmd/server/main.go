package main

import (
	"context"
	"log"
	"net"

	greetingv1 "github.com/IAmRiteshKoushik/speedrun-grpc/gen/greeting/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const listenAddr = "localhost:50051"

type greetingServer struct {
	greetingv1.UnimplementedGreeterServiceServer
}

func (greetingServer) Greet(ctx context.Context, request *greetingv1.GreetRequest) (*greetingv1.GreetResponse, error) {
	// Fixing context deadline exceeded
	// select {
	// case <-time.After(3 * time.Second):
	// 	return &greetingv1.GreetResponse{
	// 		Message: "Hello " + request.GetName(),
	// 	}, nil
	// case <-ctx.Done():
	// 	log.Printf("Greet cancelled: %v", ctx.Err())
	// 	return nil, ctx.Err()
	// }

	requestIDs := metadata.ValueFromIncomingContext(ctx, "x-request-id")
	if len(requestIDs) == 0 || requestIDs[0] == "" {
		return nil, status.Error(codes.InvalidArgument, "x-request-id is required")
	}
	log.Printf("Greet request ID: %s", requestIDs[0])

	return &greetingv1.GreetResponse{
		Message: "Hello " + request.GetName(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("listen on %s: %v", listenAddr, err)
	}

	server := grpc.NewServer()
	greetingv1.RegisterGreeterServiceServer(server, greetingServer{})

	log.Printf("gRPC server listening on %s", listenAddr)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("server gRPC: %v", err)
	}
}
