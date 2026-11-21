# Bottom-k Sketch

A bottom-k sketch retains the k smallest hash values from a set or stream.
The retained hashes form a compact sample that supports cardinality estimation.
Comparing two sketches can estimate set overlap and similarity.
Sketches built with the same hash scheme can be merged by keeping the smallest k hashes.
Increasing k improves estimates while using more memory.
Use it for approximate set analytics on large datasets.
