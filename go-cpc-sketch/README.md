# CPC Sketch

CPC (Compressed Probabilistic Counting) estimates distinct-item cardinality.
It summarizes hashed observations using compact, sparse-to-dense bit patterns.
The representation is designed to compress efficiently while retaining statistical accuracy.
Sketches support merging for distributed distinct counts.
Accuracy and storage are controlled by the sketch size parameter.
Use it for approximate unique counts when compact storage and merging matter.
