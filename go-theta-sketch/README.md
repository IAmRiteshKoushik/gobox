# Theta Sketch

A Theta sketch summarizes a set while retaining a hash-based sample.
Its theta threshold determines which hashed items are kept.
It estimates cardinality and supports set operations such as union and intersection.
Compatible sketches can be combined without revisiting the original data.
Accuracy depends on the retained sample size and overlap between sets.
Use it for approximate analytics over large or distributed datasets.
