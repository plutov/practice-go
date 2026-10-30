# Five Go exercises worth trying

[practice-go](https://github.com/plutov/practice-go) is a collection of small Go
challenges, each with a problem description, tests, and benchmarks. You can try one
on its own, then compare your approach with the contributed solutions.

I picked five that have interesting implementations. One uses Go's parser to count
function calls in source code. Another works out which nodes in a ring are broken,
even though some of their reports can't be trusted. There are also word puzzles and
a compression exercise where the output format matters as much as the algorithm.

Below, I'll walk through the parts I found worth a closer look, including a few
limitations and things to try yourself. If you'd rather solve the exercises first,
follow the challenge links before reading on.

Code snippets are excerpts, not standalone programs; adapted versions are labeled.

## 1. Function Frequency: Use Go to Understand Go

[Read the challenge](https://github.com/plutov/practice-go/tree/master/functionfrequency)
· [Read the implementation](https://github.com/plutov/practice-go/blob/master/functionfrequency/functionfrequency.go)

Given a Go source file, find its three most frequently used function names. Count
appearances in the source, not executions: a call inside a million-iteration loop
still appears only once. You're building a small static-analysis tool, not a profiler.

At first, this sounds like a job for a regular expression: find names followed by `(`.
But declarations also contain parentheses. Strings can contain text like `save()`.
Comments can describe calls that never happen. Each patch to the pattern brings you
closer to writing your own incomplete Go parser.

Go offers a more interesting starting point: its standard library already knows Go.

### Trade raw text for a syntax tree

The implementation hands the input bytes to [`go/parser`](https://pkg.go.dev/go/parser):

```go
fset := token.NewFileSet()
f, err := parser.ParseFile(fset, "src.go", gocode, 0)
if err != nil {
	return nil
}
```

Because `gocode` supplies the contents, `"src.go"` is a filename label, not a file that
must exist on disk. The file set tracks source positions, which becomes useful if
you later want diagnostics such as “this call appears on line 42.”

The result is an [abstract syntax tree](https://en.wikipedia.org/wiki/Abstract_syntax_tree),
or AST. Instead of asking which characters look like a call, you inspect nodes that
the parser has already identified as call expressions.

For the illustrative expression `fmt.Println(build())`, the relevant structure is:

```text
CallExpr
|-- Fun: SelectorExpr
|   |-- X: Ident "fmt"
|   `-- Sel: Ident "Println"
`-- Args
    `-- CallExpr
        `-- Fun: Ident "build"
```

Now strings, declarations, and calls are different kinds of data, not competing guesses
about the same sequence of characters.

### Let interfaces describe the tree; use type switches to inspect it

[`ast.Inspect`](https://pkg.go.dev/go/ast#Inspect) walks the tree and invokes a callback.
The implementation uses a type switch to select `*ast.CallExpr` nodes, then another
to distinguish a plain name such as `build` from a selector such as `fmt.Println`.

Here's a condensed version of that traversal, using `counts` for the frequency map:

```go
counts := make(map[string]int)
ast.Inspect(f, func(n ast.Node) bool {
	switch x := n.(type) {
	case *ast.CallExpr:
		switch fn := x.Fun.(type) {
		case *ast.Ident:
			counts[fn.Name]++
		case *ast.SelectorExpr:
			if receiver, ok := fn.X.(*ast.Ident); ok {
				counts[receiver.Name+"."+fn.Sel.Name]++
			}
		}
	}
	return true
})
```

It's a practical use of Go's interfaces and [type switches](https://go.dev/tour/methods/16):
a shared node interface lets the traversal handle many shapes, while concrete types
expose the fields needed for each case.

The callback returns `true` to keep descending. That detail matters: in the diagram,
it lets the traversal count both the outer `fmt.Println` and the nested `build` call.
Returning `false` at the outer call would skip its children.

### Syntax tells you the shape, not the identity

There's an important boundary here. A selector might refer to a package function or
a method on a variable. The implementation records its written name; it doesn't
resolve the declaration behind it. An import alias therefore changes the recorded name.

Also, Go represents a conversion such as `int(x)` with a `CallExpr`. The implementation
counts that shape too. It is a useful syntax-based counter, not a complete semantic
inventory of function calls; more complex callee expressions need additional handling.

Go's [`go/types`](https://pkg.go.dev/go/types) package is the next layer when you need
to distinguish conversions from calls or connect identifiers to declared objects.
That's a compelling progression: a small exercise can grow into a linter or code browser
without leaving Go's tooling ecosystem.

**What this teaches:** don't reimplement a language grammar when structured data is
already available. Go's standard library makes parsing and inspecting Go source an
ordinary programming task, not something reserved for compiler authors.

**Try next:** report each call's line number with the file set. Then add alphabetical
tie-breaking to the frequency ranking: map iteration order is unspecified, and the
current count-only sort doesn't guarantee a repeatable order for equally common names.

## 2. Build Word: Remember the Best Prefix, Not Every Attempt

[Read the challenge](https://github.com/plutov/practice-go/tree/master/buildword)
· [Read the DP implementation](https://github.com/plutov/practice-go/blob/master/buildword/buildword_dp.go)

You have a target word and a collection of reusable fragments. Build the word using
as few fragments as possible. Fragments cannot overlap, and if no construction exists,
the function returns zero.

### Why the obvious greedy approach fails

Consider the target `buildword` and these fragments:

```text
buil   dwor   bu   ild   wo   rd
```

Taking the longest matching fragment first chooses `buil`. The next fragment `dwor`
gets you almost to the end—but leaves a final `d` that no fragment can supply.

A successful construction starts with the shorter choice:

```text
bu + ild + wo + rd = buildword
```

“Make the most progress now” isn't the same as “finish with the fewest pieces.”

The [dynamic programming](https://en.wikipedia.org/wiki/Dynamic_programming) solution
uses `dp[i]` to mean the fewest fragments needed to construct `word[:i]`.

That definition is the important step. Once you've reached a prefix, the remaining
choices don't depend on the exact sequence used to reach it. Fragments are reusable,
so there is no inventory of consumed pieces to remember.

For this example, the successful path through prefix lengths looks like this:

```text
position:  0 --bu--> 2 --ild--> 5 --wo--> 7 --rd--> 9
cost:      0        1         2        3        4
```

The implementation initializes the table like this:

```go
dp := make([]int, n+1)
for i := 1; i <= n; i++ {
	dp[i] = n + 1
}
```

Go's zero initialization gives us `dp[0] == 0`: building an empty prefix costs nothing.
The other entries begin at an unreachable value. With nonempty fragments, a successful
construction cannot use more than `n` pieces, so `n + 1` is a convenient sentinel.

For each ending position and each fragment, the algorithm checks whether that fragment
can extend a reachable prefix:

```go
if dp[prevIndex] <= n && word[prevIndex:i] == fragment {
	if dp[prevIndex]+1 < dp[i] {
		dp[i] = dp[prevIndex] + 1
	}
}
```

The comparison checks the fit; the addition charges one fragment; the minimum keeps the best cost.

### A table can also be a graph

There's a second way to understand this solution: prefix lengths are vertices, and a
matching fragment creates an edge from its starting position to its ending position.
Every edge costs one fragment.

With nonempty fragments, edges always move forward. Processing prefix lengths in order
is therefore a shortest-path calculation on a directed acyclic graph. Dynamic
programming and graph algorithms aren't always separate boxes; sometimes they're two
views of the same computation.

For a target of byte length `n`, `F` fragments, and maximum fragment byte length `L`,
a straightforward upper bound is `O(nFL)` time, including string comparisons, and
`O(n)` extra space for the table.

**What this teaches:** the hard part of dynamic programming is choosing what to remember.
Here, one cost per prefix is sufficient. If fragments could only be used once, that
state would no longer capture everything needed to make the next decision.

**Try next:** store the chosen previous position alongside each cost, then reconstruct
the fragments used. Return an explanation, not just a number.

## 3. Word Ladder: Stop Searching Strings. Start Exploring a Graph.

[Read the challenge](https://github.com/plutov/practice-go/tree/master/wordladder)
· [Read the implementation](https://github.com/plutov/practice-go/blob/master/wordladder/wordladder.go)

Start with one word and reach another by changing one letter at a time. Intermediate
words must belong to the dictionary. Find the shortest transformation sequence.

Two words are neighbors if they differ at exactly one position. The implementation
compares dictionary pairs and records both directions of each connection:

```go
if dist(dic[i], dic[j]) {
	adj[i] = append(adj[i], j)
	adj[j] = append(adj[j], i)
}
```

An adjacency list stores relationships as integer indexes rather than repeatedly
copying strings into the search structure. A simplified view of two candidate routes
looks like this; additional cross-links between dictionary words are omitted:

```text
          dot --- dog
         /           \
hit --- hot           cog
         \           /
          lot --- log
```

The question is now “what is the shortest path?” rather than “which letter should I try?”

### Explore by distance, not by commitment

[Breadth-first search](https://cp-algorithms.com/graph/breadth-first-search.html), or BFS,
uses a queue to explore all vertices at one distance before moving to the next.

That ordering gives us the shortest-path guarantee. If every transition costs one,
there can't be an undiscovered shorter route hiding behind a deeper search branch.

```text
distance 0:  hit
distance 1:  hot
distance 2:  dot, lot
distance 3:  dog, log
distance 4:  cog
```

Notice the distinction between distance and the requested answer. Four transformations
visit five words. The repository returns `d + 1` when it reaches the target because
the challenge counts words in the sequence, not edges in the graph.

### The implementation leaves useful room to improve

Building all pairwise connections costs `O(N²L)` for `N` words of length `L`.
For a large dictionary, graph construction can dominate the search.

One alternative is to index wildcard patterns:

```text
hot -> *ot, h*t, ho*
dot -> *ot, d*t, do*
lot -> *ot, l*t, lo*
```

Words sharing a pattern become candidate neighbors. This avoids comparing every pair,
although large buckets still need careful handling. During BFS, processing a pattern
bucket once can avoid repeatedly scanning the same neighbor group.

There's also a smaller refinement. The current BFS marks a vertex visited when it
removes it from the queue, allowing multiple parents to enqueue that vertex first.
Marking it when enqueued prevents that duplication. The enqueue step would look like:

```go
if !vis[v] {
	vis[v] = true
	q.Add(Pair{v, d + 1})
}
```

Finally, the comparison helper assumes equal-length, byte-oriented words. Mixed lengths
need validation or filtering, and Unicode text requires a deliberate definition of
“one letter.” A byte, a Unicode code point, and a visible character aren't synonymous.

**What this teaches:** BFS is a tool for equal-cost transitions, not just textbook graphs.
It also applies to move sequences, configuration changes, and many small planning tasks.

**Try next:** record each word's predecessor and return the ladder itself. Then explore
[bidirectional search](https://en.wikipedia.org/wiki/Bidirectional_search), starting at
both endpoints rather than expanding only from the source.

## 4. Broken Nodes: Sometimes the Correct Answer Is “We Don't Know”

[Read the challenge](https://github.com/plutov/practice-go/tree/master/brokennode)
· [Read the implementation](https://github.com/plutov/practice-go/blob/master/brokennode/brokennode.go)

This is the most unusual puzzle in the group. Nodes form a ring. Each reports whether
the next node works. Working nodes report accurately; broken nodes may report anything.
You also know exactly how many nodes are broken.

The goal isn't simply to find a plausible arrangement. For each node, report whether
it must be working, must be broken, or cannot be determined from the evidence.

### A broken node isn't necessarily a liar

Suppose the three reports are:

```text
node 0 says: node 1 works
node 1 says: node 2 is broken
node 2 says: node 0 is broken
```

If a reporting node works, its statement constrains its neighbor. If it's broken,
its statement imposes no constraint—it could accidentally be true.

Unknown reliability is not reliable negation: a broken node needn't always lie.

With exactly one broken node, the reports imply `WWB`. With exactly two broken nodes,
there are two valid arrangements:

```text
possibility A:  B W B
possibility B:  B B W
                -----
what we know:   B ? ?
```

### Encode each possible world as bits

The implementation uses a `uint64` as a candidate configuration. Bit `i` represents
node `i`: `1` for broken and `0` for working.

Knowing the broken-node count means we only need masks with that many set bits. The
implementation advances directly between those masks using a bit-manipulation trick
for the [next bit permutation](https://graphics.stanford.edu/~seander/bithacks.html#NextBitPermutation).

That is more targeted than testing all `2^n` masks and counting their bits afterward.
For `n` nodes and `k` broken nodes, the candidate count is the binomial coefficient
`C(n, k)`. Checking each candidate takes a pass over the reports.

It's still a combinatorial search, especially when `k` is near half of `n`. The compact
representation makes it practical for small instances; it doesn't remove that growth.

### Make the ring look like a line

One of the nicest implementation details is this expression:

```go
c := conf | ((conf & 1) << uint8(len(reports)))
```

`conf & 1` extracts the first node's status. Shifting it by the number of nodes puts a
copy immediately beyond the last node. Now shifting through the configuration can
compare every node with its successor, including the final-to-first pair.

The representation handles the boundary case, keeping it out of the checking loop.

This trick also makes the fixed-width boundary important. The wraparound copy needs
an extra bit, so this isn't an implementation for arbitrary ring sizes; zero, maximum
width, and inconsistent-input cases deserve explicit contracts and tests.

### Merge possibilities without inventing certainty

After checking a valid configuration, the implementation merges it into the result:

```go
switch {
case result[i] == 0:
	result[i] = chr(conf)
case result[i] != chr(conf):
	result[i] = '?'
}
```

The first valid configuration initializes each status. A disagreement changes it to
`?`. Once a position becomes uncertain, later possibilities cannot make it certain
again: the earlier disagreement remains evidence.

**What this teaches:** satisfying constraints and drawing conclusions from constraints
are different tasks. This is a small, approachable example of reasoning across all
valid models, rather than trusting the first model you find.

**Try next:** also return the valid-configuration count to distinguish no solution from
one explanation or several. For context, read about [constraint satisfaction](https://en.wikipedia.org/wiki/Constraint_satisfaction_problem).

## 5. Compression: Short Codes Are Not the Same as Small Files

[Read the challenge](https://github.com/plutov/practice-go/tree/master/compression)
· [Read the implementation](https://github.com/plutov/practice-go/blob/master/compression/compression.go)

Encode a string, decode it, and get the original back. This exercise includes a sample
dataset and benchmark for connecting an algorithm to output size and execution cost.

The submitted implementation uses [Huffman coding](https://en.wikipedia.org/wiki/Huffman_coding).
Its central idea is intuitive: frequent symbols should receive short codes, while
infrequent symbols can afford longer ones.

### Build the code from the data

The implementation counts runes, places their frequencies into a minimum priority
queue, and repeatedly combines the two least-frequent nodes.

For illustrative frequencies `A:5`, `B:2`, and `C:1`, one possible tree is:

```text
                  (8)
                 /   \
              0 /     \ 1
               A:5    (3)
                     /   \
                  0 /     \ 1
                   C:1    B:2

              A = 0
              C = 10
              B = 11
```

Left and right assignments can differ without changing the idea. In this tree, eight
symbols need eleven payload bits, compared with sixteen using fixed two-bit codes.
That comparison excludes the codebook, which we'll come back to in a moment.

The tree-building merge is small:

```go
c := &HNode{
	freq:  a.freq + b.freq,
	left:  a,
	right: b,
}
hfNodesPQ.Insert(c)
```

The repository supplies its own [minimum-priority-queue implementation](https://github.com/plutov/practice-go/blob/master/compression/minpq.go),
using a binary heap. With `U` distinct symbols, heap-based tree construction takes
`O(U log U)` time. Go's standard library also offers [`container/heap`](https://pkg.go.dev/container/heap)
if you'd like to compare the interfaces and implementation choices.

### Why separators aren't necessary

Huffman codes are [prefix-free](https://en.wikipedia.org/wiki/Prefix_code): no complete
symbol code is the prefix of another. In the example, seeing `0` identifies `A`
immediately. Seeing `1` tells the decoder to read another bit.

The tree has moved symbol-boundary information into the code itself. No delimiters needed.

For a single distinct rune, the implementation assigns `"0"` rather than an empty code.
Otherwise, writing that symbol repeatedly would produce no payload at all.

### The revealing detail: these bits are still bytes

The implementation builds codes as strings of `'0'` and `'1'`. Then it writes them:

```go
encodedBytesBuffer.WriteString(charCodeMap[el])
```

This demonstrates Huffman coding, but those characters each occupy a full byte.
An eight-digit code string is eight bytes, not one packed byte. The codebook is also
serialized alongside the payload, adding overhead that the decoder needs to recover
the original symbols.

So a successful round trip doesn't prove successful compression. The current test
checks that decoding restores the sample and prints an output-size ratio; it doesn't
require the encoded representation to be smaller.

### Finish the format, not just the algorithm

A bit-packed version needs to accumulate bits into bytes and track how many bits in
the last byte are meaningful. Otherwise, padding zeros can be mistaken for more data.

One possible format—not the repository's current format—would be:

```text
+-----------------------+
| Codebook length       |
+-----------------------+
| Serialized codebook   |
+-----------------------+
| Valid payload bits    |
+-----------------------+
| Packed payload bytes  |
+-----------------------+
```

For tiny inputs, even a packed payload may not offset the header and codebook.
Measure the complete encoded output, not just the interesting part.

There's also a Go-specific contract to settle: the implementation operates on runes,
while a Go string can contain arbitrary bytes. A general-purpose byte compressor
should preserve invalid UTF-8 too; converting through runes isn't sufficient for that.
See [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) for the distinction.

**What this teaches:** lossless decoding, compact representation, and robust file-format
design are related but separate goals. Tests should make those goals explicit.

**Try next:** pack the payload into bytes, test empty and single-symbol inputs, and
compare total output size on repetitive text, ordinary prose, and poorly compressible
data. Don't expect one technique to shrink every input.

## How to Get More Than Five Solutions Out of These Exercises

A concise solution can make an idea seem obvious in hindsight. Close the source file
and recreate the reasoning instead. For each exercise:

1. Write down the state: a syntax node, prefix position, graph vertex, bitmask, or code tree.
2. Explain what each transition preserves and why it moves toward an answer.
3. Implement the simplest correct version you can explain without hand-waving.
4. Add an edge case that challenges one of your assumptions.
5. Measure before introducing a more complicated optimization.

From the repository root, a focused test-and-benchmark run looks like this:

```sh
go test ./buildword -bench . -benchmem
```

See the [`testing` documentation](https://pkg.go.dev/testing#hdr-Benchmarks) for benchmark details.
Compare implementations on the same machine with representative inputs.

Across these five exercises, the reusable insight is not a particular clever line.
It's the choice of what to keep, what to ignore, and what to make explicit:

- Function Frequency replaces text-matching guesses with Go's own syntax tree.
- Build Word keeps the cheapest prefix instead of every construction history.
- Word Ladder replaces letter guessing with equal-cost graph transitions.
- Broken Nodes combines possible worlds without pretending uncertainty disappeared.
- Compression makes symbol boundaries implicit, then forces us to account for bytes.

Small problems give you room to inspect decisions that disappear inside large applications.
Pick one challenge, sketch its representation, and explain it before optimizing.
Then ask: **what becomes simpler if I change the shape of the problem?**
