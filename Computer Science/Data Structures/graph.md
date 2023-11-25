# Graph
Graphs are a kind of [data structure](Computer%20Science/Data%20Structures/data%20structure.md) that are made up of nodes (also called vertices) that connect to each other through lines. We call these lines "edges". Quite similar to the [tree](Computer%20Science/Data%20Structures/tree.md), really. The difference is that the graph does not *require* a hierarchical structure.

When the edges in a graph have a direction, we call it a [directed graph](Computer%20Science/Data%20Structures/Graph/directed%20graph.md).  
When they do not have a specified direction, we call it an [undirected graph](Computer%20Science/Data%20Structures/Graph/undirected%20graph.md).  
The edges may also have a *weight* or a *value* associated to them, leading to another type of structure: the weighted graph. If a graph is unweighted, it is assumed that all edges have the same weight.

Other relevant classifications of graphs:
- Cyclic graph: Contains at least one cycle, i.e. a path that begins and ends in the same node.
- Acyclic graph and Directed Acyclic Graph (DAG): The first is a graph with no cycles -- the second would be a graph with directions, but no paths looping back to the same node/vertex. These are often used in scheduling, task dependencies, and representing hierarchical structures.
- Connected graph: When there is a path between every pair of nodes.
- Complete graph: When there is an edge between every pair of nodes.
- Planar graph: When the graph can be inserted in a 2D plane with no edges intersecting.
- Regular graph: When all nodes have the same degree (i.e. the same number of edges connected to them)
## Sources
- [Graph Data Structure](https://www.simplilearn.com/tutorials/data-structure-tutorial/graphs-in-data-structure)
- [Graph Data Structure | Illustrated Data Structures](https://www.youtube.com/watch?v=0sQE8zKhad0)