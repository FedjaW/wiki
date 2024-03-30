# Graphs

- All trees are graphs, not all graphs are trees

## Terminology of Graphs

**Graph Terms:**

This is not an exhaustive list of terms, but it is the terms that we may end up using today.

- cycle: When you start at Node(x), follow the links, and end back at Node(x)
- acyclic: A graph that contains no cycles
- connected: When every node has a path to another node
- directed: When there is a direction to the connections. Think Twitter
- undirected: !directed.
- weighted: The edges have a weight associated with them. Think Maps
- dag: Directed, acyclic graph.

**Implementation Terms:**

- node: a point or vertex on the graph
- edge: the connection betxit two nodes

**Big O:**

- BigO is commonly stated in terms of V and E where V stands for vertices and E stands for edges
- So O(V \* E) means that we will check every vertex, and on every vertex we check every edge
