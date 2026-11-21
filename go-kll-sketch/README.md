# KLL Sketch

KLL is a compact, mergeable sketch for approximate quantile queries.
It keeps weighted samples in levels and compacts them as capacity fills.
The result estimates ranks and quantiles without retaining every observation.
Its memory use is controlled by a configurable accuracy parameter.
Use it for streaming latency, measurement, or distribution summaries.
Quantiles are approximate; the sketch does not provide exact order statistics.
