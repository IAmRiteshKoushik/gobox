# Greenwald–Khanna (GK) Sketch

The GK sketch estimates quantiles in a streaming sequence of numbers.
It stores a compact summary of values and rank bounds rather than the full stream.
Its deterministic rank error is controlled by a chosen epsilon parameter.
Insertions update the summary, which periodically compresses redundant entries.
Use it when a bounded rank guarantee matters and data arrives incrementally.
The returned quantile value is approximate, with rank error bounded by the configuration.
