# Bloom Filter

A Bloom filter answers approximate set-membership queries.
Several hashes mark bit positions when an item is inserted.
Queries can return false positives, but a correctly configured filter has no false negatives.
The bit array is compact and supports fast insertion and lookup.
Capacity and target false-positive rate determine its memory and hash count.
Use it to avoid expensive lookups when an item is probably absent.
