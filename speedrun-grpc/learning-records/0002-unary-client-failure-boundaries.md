# Unary client failure boundaries

The learner generated a Go gRPC client from a `.proto` contract and wrote a compiling unary call that owns and closes its connection, bounds the RPC with `context.WithTimeout`, and classifies status errors. This establishes a practical base for a successful local call and later metadata or streaming exercises.

## Evidence

`cmd/client/main.go` creates `GreeterServiceClient`, calls `Greet`, and branches on `Unavailable` and `DeadlineExceeded`; `mise exec -- go test ./...` passes.
