# Status Codes in gRPC

| gRPC code | Name                  | Rough HTTP analogue |
| --------: | --------------------- | ------------------- |
|         0 | `OK`                  | 200                 |
|         1 | `CANCELLED`           | 499-ish             |
|         2 | `UNKNOWN`             | 500                 |
|         3 | `INVALID_ARGUMENT`    | 400                 |
|         4 | `DEADLINE_EXCEEDED`   | 504                 |
|         5 | `NOT_FOUND`           | 404                 |
|         6 | `ALREADY_EXISTS`      | 409                 |
|         7 | `PERMISSION_DENIED`   | 403                 |
|         8 | `RESOURCE_EXHAUSTED`  | 429                 |
|         9 | `FAILED_PRECONDITION` | 400/409             |
|        10 | `ABORTED`             | 409                 |
|        11 | `OUT_OF_RANGE`        | 400                 |
|        12 | `UNIMPLEMENTED`       | 501                 |
|        13 | `INTERNAL`            | 500                 |
|        14 | `UNAVAILABLE`         | 503                 |
|        15 | `DATA_LOSS`           | 500                 |
|        16 | `UNAUTHENTICATED`     | 401                 |
