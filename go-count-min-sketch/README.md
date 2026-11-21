# Count-Min Sketch

Count-Min Sketch estimates how often keys occur in a stream.
It hashes each key into several rows of a compact counter table.
An estimate is the minimum counter among that key's hashed positions.
This gives one-sided estimates: collisions can overcount, but not undercount.
Choose width and depth to trade memory for error and confidence.
It is useful for heavy-hitter screening and approximate frequency queries.
