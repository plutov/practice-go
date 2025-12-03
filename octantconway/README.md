### 𜴂𜴯 Octant Conway 𜴩🯦

Conway's Game of Life (GoL) is a cellular automaton consisting of a (possibly infinite) grid of either live or dead square cells. The cell arrangement changes over time. Cells are influenced solely by their eight surrounding neighbors. At each step cell state reevaluates according to following rules:
- **L->D n<2**: a live cell with fewer than two live neighbors dies
- **L->D n>3**: a live cell with more than three live neighbors dies
- **D->L n=3**: a dead cell with exactly three live neighbors becomes alive

Please write a function `OctantConway([]byte) []byte` to compute the next generation of a given GoL configuration. The configuration is provided as a slice of bytes containing UTF-8 encoded text. Cells are represented mainly through Octant Unicode characters. Each character contains eight (2x4) cells. Return the result accordingly.

> Note: the Octant Unicode block does not contain all needed characters (256) to avoid duplication. Some shared characters are located elsewhere.
> - Octant Unicode characters (230): u+1cd00...u+1cde5
> - Quadrant Unicode characters (10): u+2596...u+259f
> - Other characters (16): space ' ', full block '█', ...

### Run tests

Run animated tests that show the evolution of the grid:

```
go test -run='TestOctantConway'
```

A single test runs around one second at maximum. Skip animations and run tests as fast as possible with:

```
go test -run='TestOctantConway' -short
```

> Note: `go test .` doesn't show animations as it buffers output until test completion.
