# DDSketch

DDSketch is a mergeable sketch for approximate distribution quantiles.
It groups values into logarithmic bins, keeping separate bins for positive and negative values.
The logarithmic mapping provides a relative-error guarantee for represented quantiles.
Memory can be bounded by collapsing bins when the configured limit is reached.
It is well suited to latency and other measurements spanning many magnitudes.
Quantile results are estimates, not exact values from the original stream.
