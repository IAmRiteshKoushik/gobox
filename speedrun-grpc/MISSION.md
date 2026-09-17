# Mission: become productive with a Go gRPC client

## Why

Tomorrow, I need to create and work with a gRPC client in a Go codebase. I have more than two years of Go experience, but no practical gRPC experience. I need enough command of the client path to read the contract, make safe calls, interpret failures, and ask precise questions when repository conventions are unclear.

## Success looks like

- I can find the `.proto` contract and identify the generated client, request, response, and RPC shape.
- I can create or use a `grpc.ClientConn`, call a unary RPC with a deadline, and close owned resources.
- I can diagnose status errors, attach outgoing metadata when the service requires it, and distinguish local setup failures from RPC failures.
- I can follow my team's existing code-generation, credentials, observability, and retry conventions without inventing a parallel client stack.

## Constraints

- Three to four hours today, focused on client work.
- Strong Go fluency. Skip language and basic concurrency instruction.
- The local shell uses `mise exec -- go ...`; `protoc` exists and `buf` does not.

## Out of scope

- Designing a production server.
- Owning a protobuf schema or compatibility policy beyond reading one safely.
- Load balancing, service mesh internals, and advanced authentication unless the work repository needs them.
