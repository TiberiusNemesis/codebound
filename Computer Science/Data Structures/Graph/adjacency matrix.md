# Adjacency Matrix

A 2D array representation of a graph, where the eleemnts in the matrix represent the adjacency (or lack of it) in vertice pairs. It is a straightforward representation that applies to both the directed graph and the undirected graph.

The matrix has an N x N format, N being the number of vertices.

An element `M[i][j]` in the matrix is equal to 1 (or a weight number, if it is a weighted graph) only if there is an edge from vertex i to j. Otherwise, if there is no edge, it is 0.

Example:

```c
  A B C D
A 0 1 1 0
B 1 0 0 1
C 1 0 0 1
D 0 1 1 0
```

In this graph, we have the following connections:

A to B, and C  
B to A, and D  
C to A, and D  
D to B, and C  

The main advantages of this type of graph representation is that it is quite fast at checking if an edge exists between any two vertices, and it is very simple and straightforward.

## Sources

- [Adjacency Matrix - Graph Representation](https://www.programiz.com/dsa/graph-adjacency-matrix)
