# Misra–Gries

Misra–Gries finds frequent items in a stream with a bounded set of counters.
Known items increment their counters; new items use free counters when available.
When all counters are occupied, a pass decrements them and removes zeros.
Any item above a frequency threshold is guaranteed to appear among the candidates.
The algorithm uses deterministic memory proportional to the candidate limit.
Use it to identify heavy hitters, then optionally count candidates in a second pass.
