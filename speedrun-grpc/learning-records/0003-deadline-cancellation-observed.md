# Deadline cancellation observed end to end

The learner ran a unary RPC whose server delay exceeded the client's two-second budget and observed `DeadlineExceeded` at the client plus `context canceled` through the server request context. This confirms that a deadline must be paired with server-side cancellation-aware work to release resources.

## Evidence

The client used `context.WithTimeout`; the server selected between delayed work and `ctx.Done()` and logged the cancellation outcome.
