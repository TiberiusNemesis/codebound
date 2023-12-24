# Adjacency Matrix

A 2D [array](Computer%20Science/Data%20Structures/array.md) representation of a [graph](Computer%20Science/Data%20Structures/graph.md), where the eleemnts in the matrix represent the adjacency (or lack of it) in vertice pairs. It is a straightforward representation that applies to both the [directed graph](Computer%20Science/Data%20Structures/Graph/directed%20graph.md) and the [undirected graph](Computer%20Science/Data%20Structures/Graph/undirected%20graph.md).

The matrix has an N x N format, N being the number of vertices.

An element `M[i][j]` in the matrix is equal to 1 (or a weight number, if it is a weighted [graph](Computer%20Science/Data%20Structures/graph.md)) only if there is an edge from vertex i to j. Otherwise, if there is no edge, it is 0.

Example:

```c
  A B C D
A 0 1 1 0
B 1 0 0 1
C 1 0 0 1
D 0 1 1 0
```

In this [graph](Computer%20Science/Data%20Structures/graph.md), we have the following connections:

A to B, and C  
B to A, and D  
C to A, and D  
D to B, and C  

The main advantages of this type of [graph](Computer%20Science/Data%20Structures/graph.md) representation is that it is quite fast at checking if an edge exists between any two vertices, and it is very simple and straightforward.

## Sources

- [Adjacency Matrix - Graph Representation](https://www.programiz.com/dsa/graph-adjacency-matrix)
