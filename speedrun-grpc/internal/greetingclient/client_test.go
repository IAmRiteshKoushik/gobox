package greetingclient_test

import (
	"context"
	"fmt"
	"net"
	"slices"
	"sync"
	"testing"
	"time"

	greetingv1 "github.com/IAmRiteshKoushik/speedrun-grpc/gen/greeting/v1"
	"github.com/IAmRiteshKoushik/speedrun-grpc/internal/greetingclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type testGreeterServer struct {
	greetingv1.UnimplementedGreeterServiceServer

	interruptAfterFirst bool

	mu        sync.Mutex
	requestID string
}

func (s *testGreeterServer) GreetMany(
	request *greetingv1.GreetRequest,
	stream grpc.ServerStreamingServer[greetingv1.GreetResponse],
) error {
	requestIDs := metadata.ValueFromIncomingContext(stream.Context(), "x-request-id")
	if len(requestIDs) == 0 || requestIDs[0] == "" {
		return status.Error(codes.InvalidArgument, "x-request-id is required")
	}

	s.mu.Lock()
	s.requestID = requestIDs[0]
	s.mu.Unlock()

	for i := range 3 {
		if s.interruptAfterFirst && i == 1 {
			return status.Error(
				codes.Unavailable,
				"simulated stream interruption",
			)
		}

		if err := stream.Send(&greetingv1.GreetResponse{
			Message: fmt.Sprintf(
				"Hello %s #%d",
				request.GetName(),
				i+1,
			),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *testGreeterServer) receivedRequestID() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.requestID
}

func newTestClient(t *testing.T, server greetingv1.GreeterServiceServer) *greetingclient.Client {
	t.Helper()

	listener := bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()
	greetingv1.RegisterGreeterServiceServer(grpcServer, server)

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("create test client connection: %v", err)
	}

	t.Cleanup(func() {
		_ = conn.Close()
	})

	return greetingclient.New(conn)
}

func TestGreetManyReceivesEveryMessage(t *testing.T) {
	server := &testGreeterServer{}
	client := newTestClient(t, server)

	ctx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()

	messages, err := client.GreetMany(
		ctx,
		"Ritesh",
		"test-request-001",
	)
	if err != nil {
		t.Fatalf("GreetMany returned an error: %v", err)
	}

	want := []string{
		"Hello Ritesh #1",
		"Hello Ritesh #2",
		"Hello Ritesh #3",
	}
	if !slices.Equal(messages, want) {
		t.Fatalf("messages = %q, want %q", messages, want)
	}

	if got := server.receivedRequestID(); got != "test-request-001" {
		t.Fatalf("request ID = %q, want %q", got, "test-request-001")
	}
}

func TestGreetManyPreservesUnavailable(t *testing.T) {
	server := &testGreeterServer{
		interruptAfterFirst: true,
	}
	client := newTestClient(t, server)

	_, err := client.GreetMany(
		t.Context(),
		"Ritesh",
		"test-request-002",
	)
	if got := status.Code(err); got != codes.Unavailable {
		t.Fatalf("status code = %s, want %s; error = %v", got, codes.Unavailable, err)
	}
}

func TestGreetManyRejectsMissingRequestID(t *testing.T) {
	server := &testGreeterServer{}
	client := newTestClient(t, server)

	_, err := client.GreetMany(t.Context(), "Ritesh", "")
	if got := status.Code(err); got != codes.InvalidArgument {
		t.Fatalf(
			"status code = %s, want %s; error = %v",
			got,
			codes.InvalidArgument,
			err,
		)
	}
}
