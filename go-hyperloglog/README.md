# HyperLogLog

HyperLogLog estimates the number of distinct values in a stream.
It hashes each value and records leading-zero patterns in a small register array.
Register statistics combine into a cardinality estimate with near-constant space.
More registers improve accuracy, and independent sketches can be merged.
Use it for approximate unique-user, event, or key counts at large scale.
The estimate is probabilistic and does not identify the distinct values.
