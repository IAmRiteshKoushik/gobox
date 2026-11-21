# SimHash

SimHash maps a weighted feature set to a compact fingerprint.
Feature hashes vote for fingerprint bits, with weights reflecting feature importance.
Similar inputs tend to produce fingerprints with small Hamming distance.
It is a similarity-oriented fingerprint, not an exact membership structure.
Feature extraction and bit length determine how well the fingerprint reflects similarity.
Use it for near-duplicate documents, pages, or other feature-based records.
