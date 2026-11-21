# Frequent Directions

Frequent Directions is a deterministic sketch for large matrices and data streams.
It compresses rows into a smaller matrix while preserving important covariance structure.
The sketch approximates principal components and matrix products with bounded storage.
When its buffer fills, singular values are shrunk to make room for new observations.
Its error depends on sketch size and the spectral tail of the input matrix.
Use it for streaming dimensionality reduction and approximate PCA.
