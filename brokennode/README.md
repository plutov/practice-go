### Find Broken Node

We have a chain of nodes, the functionality of each node is to find if next node is working well. Unfortunately, we recently found that some of our nodes are broken.

Each node checked next node and we have a report about it. We can totally trust reports made by working nodes, but broken nodes could return wrong results, so we can't trust them.

In the report `true` means that node was reported to be working, `false` means that node reported to be broken. Last report element is report from last node about the first node.

We also know the number of broken nodes in the system - `brokenNodes`.

Your code should return a string where every character describes the alignment of the corresponding node: `W` for working, `B` for broken, and `?` for unknown.

### Example

```
FindBrokenNodes(1, []bool{true, false, false}) // WWB
```

This is what each node reported:

- Node 1: Node 2 is working
- Node 2: Node 3 is broken
- Node 3: Node 1 is broken

So there are 4 possibilities:

- Working - Working - Broken (1 broken)
- Broken - Working - Broken (2 broken)
- Broken - Broken - Working (2 broken)
- Broken - Broken - Broken (3 broken)

In this example we know that there is exactly 1 broken node, the output should be "WWB".

```
FindBrokenNodes(2, []bool{true, false, false}) // B??
```

From the same possibilities we keep only:

- Broken - Working - Broken (2 broken)
- Broken - Broken - Working (2 broken)

In both cases, the 0th node is definitely broken, but we're not sure about the other two, so the output should be "B??".

### Run tests with benchmarks

```
go test -bench .
```

### Submitted solutions

shogg [solution](https://github.com/plutov/practice-go/blob/e92e4a814ab3d55ad9ab57392d942e50164f0384/brokennode/brokennode.go)
```
goos: linux
goarch: amd64
pkg: github.com/plutov/practice-go/brokennode
BenchmarkFindBrokenNodes-2   	20000000	        93.4 ns/op	       6 B/op	       2 allocs/op
PASS
ok  	github.com/plutov/practice-go/brokennode	1.979s
```