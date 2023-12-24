# Adjacency List

The adjacency list is a way to represent a [graph](Computer%20Science/Data%20Structures/graph.md) as a collection of lists.

Each list describes all the neighbors of a vertex in a [graph](Computer%20Science/Data%20Structures/graph.md). This representation is quite efficient in terms of space when dealing with sparse graphs (i.e. graphs where the number of edges is much less than the number of vertices squared).

For an [undirected graph](Computer%20Science/Data%20Structures/Graph/undirected%20graph.md):

Each element of the adjacency list contains a list of vertices adjacent to the element.

For a [directed graph](Computer%20Science/Data%20Structures/Graph/directed%20graph.md):

Each element of the adjacency list contains a list of vertices that the element *points to* (i.e. outgoing edges).

As an example, if we have a [graph](Computer%20Science/Data%20Structures/graph.md) with vertices A, B, C, D and edges AB, AC, BD, and CD, the adjacency list may look like:

A: `[B, C]`  
B: `[A, D]`  
C: `[A, D]`  
D: `[B, C]`  

## Sources

- [Adjacency List - Graph Representation](https://www.programiz.com/dsa/graph-adjacency-list)
