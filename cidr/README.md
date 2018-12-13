### Inverse CIDR whitelist

Write a func that turns a list of IPv6 addresses into a shortest possible list of CIDR masks, such that everything but those IPs is matched. And returns the number of masks.

[CIDR masks](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)

#### Input

Input is a list of IP addresses in full form: eight 4-digit hexadecimal groups separated by `:`.

For example: `0000:0000:0000:0000:0000:0000:0000:0000`.

#### Output

Output is the size of shortest list of CIDR masks that cover all IP addresses except those listed in input.

### Example

```go
NumberOfMasks([]string{"0000:0000:0000:0000:0000:0000:0000:0000"}) // 128
NumberOfMasks([]string{"0000:0000:0000:0000:0000:0000:0000:0000", "0000:0000:0000:0000:0000:0000:0000:0001"}) // 127
```

### Run tests with benchmarks

```
go test -bench .
```
