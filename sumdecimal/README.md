### Sum Decimal

You are given a number `n`. Your task is to return the sum of the first 1000 decimal places of the square root of `n`. The most difficult part in Go is to get the `Sqrt` with arbitrary precision.

### Example

The square root of `2` equals `1.4142135623...`, so the answer is calculated as `4 + 1 + 4 + 2 + 1 + ...`, 1000 digits altogether equals 4482.

```
SumDecimal(2) = 4482
```

### Run tests with benchmarks

```
go test -bench .
```
