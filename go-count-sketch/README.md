# Count Sketch

Count Sketch estimates item frequencies in a stream using randomized hashing.
Each update adds a signed count to one bucket in each hash table row.
An estimate combines the signed bucket values, commonly using their median.
Unlike Count-Min Sketch, its estimates can be above or below the true count.
Width and depth tune the error and confidence for a chosen memory budget.
Use it for approximate frequency queries and heavy-hitter discovery.
