# Outgoing metadata observed by the server

The learner attached `x-request-id` outgoing metadata to a deadline context, observed it at the gRPC server, and observed `InvalidArgument` when it was absent. This distinguishes typed business input from RPC transport context and gives a base for reading repository tracing and authentication patterns.

## Evidence

The local `Greet` handler logged `practice-001` from incoming metadata and returned a deliberate status error when the client omitted it.
