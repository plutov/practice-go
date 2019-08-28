### Node Degree

The degree of a node in the graph is the number of edges connected to the node. In this challenge the `Degree` function needs to calculate the degree of given node in the graph.

### Input

Function `Degree` expects 2 arguments:

- Number of nodes in the undirected graph
- Graph representation - bidemensional array of `a` and `b` pairs, showing that those two nodes are connected.
- Node number to calculate the degree.

If node is not found in the graph, the `Degree` func should return an error.

### Example

```
nodes := 3
graph := [][2]int{
    {1, 2},
    {1, 3},
}
Degree(nodes, graph, 1) // 2, err=nil
Degree(nodes, graph, 2) // 1, err=nil
Degree(nodes, graph, 3) // 1, err=nil
Degree(nodes, graph, 4) // 0, err=node 4 not found in the graph
```

### Run tests with benchmarks

```
go test -bench .
```
