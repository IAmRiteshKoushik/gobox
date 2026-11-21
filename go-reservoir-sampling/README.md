# Reservoir Sampling

Reservoir sampling maintains a uniform sample of a stream with a fixed-size buffer.
The first k items fill the reservoir; each later item gets a shrinking chance of replacing one.
After processing, every item seen so far has the same probability of being retained.
It works when the stream length is unknown and items cannot be revisited.
The reservoir size controls sample size and memory use.
Use it to inspect or analyze a representative sample of streaming data.
