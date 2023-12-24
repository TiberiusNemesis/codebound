# Spanning Tree

A spanning tree is a subgraph of an undirected graph where all the original vertices are included, but with the minimum number of edges necessary to connect them. This is usually *n-1* edges for *n* vertices.

Spanning trees do not contain any cycles. This is natural, since a cycle would not contribute to the minimal number of connections between vertices.

Note that a graph might (and most likely will) have multiple spanning tree representations, since there are multiple ways that the same number of edges can be formed.

![Image representing spanning trees for a graph G](/Assets/spanning-tree.png)

Spanning trees are widely used in network planning, routing protocols and in cluster analysis. We can think of a city's network as a huge graph -- therefore, plans can be made to deploy internet/telephone lines in a way where we use the minimum amount of lines (edges) to connect all of the locations in the city (i.e. the vertices).

## Minimum Spanning Tree

If we are working with a weighted graph, we can build a *minimum spanning tree* (or MST) from it. This is essentially a spanning tree (i.e. connects all vertices) with the minimum total summed up weight from all edges. This is quite useful because it is a direct metric for the minimum-value scenario from whatever we are measuring, such as distance or traffic.

## Sources

- [wiki, Spanning tree](https://en.wikipedia.org/wiki/Spanning_tree)
- [TutorialsPoint, Spanning Tree](https://www.tutorialspoint.com/data_structures_algorithms/spanning_tree.htm)
- [wiki, Minimum spanning tree](https://en.wikipedia.org/wiki/Minimum_spanning_tree)
