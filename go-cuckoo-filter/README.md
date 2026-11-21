# Cuckoo Filter

A Cuckoo filter approximates set membership using short item fingerprints.
Fingerprints live in buckets and can be moved during insertion to make room.
Queries may have false positives, while stored items are not missed.
Unlike a basic Bloom filter, it can support deletion of inserted items.
Bucket size and fingerprint width control capacity, memory, and error rate.
It is useful for compact, mutable membership indexes.
