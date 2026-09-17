# Go gRPC client resources

## Knowledge

- [gRPC-Go basics tutorial](https://grpc.io/docs/languages/go/basics/)
  Official end-to-end reference for generated stubs, `grpc.NewClient`, unary RPCs, and the four RPC shapes. Use it when connecting the `.proto` contract to Go client code.
- [grpc package reference](https://pkg.go.dev/google.golang.org/grpc)
  Current Go API reference. Use it to check `grpc.NewClient`, connection behavior, client interceptors, and current deprecations. `Dial` and `DialContext` are deprecated in favor of `NewClient`.
- [gRPC deadlines guide](https://grpc.io/docs/guides/deadlines/)
  Official rationale and behavior for time-bounding RPCs. Use it when choosing and propagating call time budgets.
- [gRPC status codes guide](https://grpc.io/docs/guides/status-codes/)
  Canonical meanings of status codes. Use it to decide whether a caller can recover, retry, alter its input, or surface an error.
- [Protocol Buffers Go generated-code guide](https://protobuf.dev/reference/go/go-generated/)
  Official guide to the Go message API. Use it when reading generated messages and their getters.

## Wisdom

- [grpc-go GitHub Discussions](https://github.com/grpc/grpc-go/discussions)
  Maintainer-hosted place for Go-specific behavior that documentation does not answer. Use it only after reducing a question to a small reproducible client example.

## Gaps

- The work repository's contract source, code-generation command, credential model, metadata keys, error mapping, telemetry, and retry policy are unknown. Inspect existing client calls before writing a new one.
