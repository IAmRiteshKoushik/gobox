# XOR Filter

An XOR filter is a compact probabilistic structure for static set membership.
It stores fingerprints arranged so a query combines a small number with XOR.
Queries have a configurable false-positive probability and no false negatives for built keys.
Construction is typically performed in batches; adding keys requires rebuilding.
It often offers fast lookups and low space for immutable key sets.
Use it for read-heavy indexes such as blocked-key or URL sets.
