# t-digest

A t-digest summarizes numeric data for approximate quantile estimation.
It clusters nearby observations into weighted centroids and merges those centroids.
Centroids are kept finer near distribution tails to improve extreme-quantile accuracy.
Digests can be combined, making them useful for distributed aggregation.
Compression controls the size and the accuracy trade-off.
Use it for percentile dashboards, especially tail latency measurements.
