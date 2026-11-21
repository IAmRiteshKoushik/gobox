# MinHash

MinHash creates compact signatures that approximate set similarity.
For each hash function, it retains the smallest hash value among a set's members.
The fraction of matching signature positions estimates Jaccard similarity.
Signatures can be compared without storing or exchanging the full sets.
More hash values improve accuracy at the cost of memory and comparison work.
Use it for near-duplicate detection, document similarity, and set comparisons.
