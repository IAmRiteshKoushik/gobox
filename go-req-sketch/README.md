# REQ Sketch

REQ (Relative Error Quantiles) is a mergeable sketch for approximate quantiles.
It uses compaction to retain a weighted sample of a numeric stream.
Its design focuses on relative rank accuracy, particularly near selected distribution tails.
A configuration chooses which tail receives stronger accuracy guarantees.
Use it for percentile monitoring when the tail behavior matters.
Results are approximate and depend on the sketch's accuracy and merge settings.
