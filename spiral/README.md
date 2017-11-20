### Spiral

Write a program that will display a `spiral` of `n × n` numbers.
Make sure your program uses constant (or linear) space. This means, it is not allowed to build an array before printing it (or to build another data structure consuming space with `O(pow n)`.

Sperate numbers with one space.

### Example

For example, here’s what the spiral looks like for n = 10:

```
99 98 97 96 95 94 93 92 91 90
64 63 62 61 60 59 58 57 56 89
65 36 35 34 33 32 31 30 55 88
66 37 16 15 14 13 12 29 54 87
67 38 17  4  3  2 11 28 53 86
68 39 18  5  0  1 10 27 52 85
69 40 19  6  7  8  9 26 51 84
70 41 20 21 22 23 24 25 50 83
71 42 43 44 45 46 47 48 49 82
72 73 74 75 76 77 78 79 80 81
```

### Run it

As this function doesn't have return value, but yields an output, use `go run` to test output:

```
go run spiral.go
```
