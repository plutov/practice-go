### Calculator

Create a calculator that can evaluate simple mathematical expressions. The calculator should be able to handle the basic arithmetic operations: addition, subtraction, multiplication, and division.

### Function

```go
Eval(expr string) (float64, error)
```

### Examples

```go
result, err := Eval("1 + 2")
fmt.Println(result) // 3

result, err := Eval("2 * 3")
fmt.Println(result) // 6

result, err := Eval("10 / 2 + 6")
fmt.Println(result) // 11

result, err := Eval("( 2 + 3 ) * 4")
fmt.Println(result) // 20
```

### Run tests with benchmarks

```
go test -bench .
```
