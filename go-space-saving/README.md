# Space-Saving

Space-Saving tracks frequent items using a fixed number of counters.
When a new item arrives and no counter is free, it replaces a least-counted item.
The replacement count records uncertainty, giving useful bounds on true frequency.
It is designed to find heavy hitters in streams with bounded memory.
Memory depends on the number of tracked candidates, not total stream length.
Use it to surface the most common keys, terms, or events.
