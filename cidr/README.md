### Inverse CIDR whitelist

Write a func that turns a list of IPv6 addresses into a shortest possible list of CIDR masks, such that everything but those IPs is matched. And returns the number of masks.

#### CIDR masks

A CIDR mask matches IPs by a prefix. `::/19` matches all IPs whose first 19 bits are zero, i.e. `0:0000-1fff:*:*:*:*:*:*`. Similarly, `55aa::/16` matches all IPs that start with `0101010110101010` bit sequence, i.e. `55aa:*:*:*:*:*:*:*`. A /128 always matches only one IP. A /0 matches all IPs.

#### Input

Input is a list of IP addresses in full form: eight 4-digit hexadecimal groups separated by : in network order (most significant first). For example, `0000:0000:0000:0000:0000:0000:0000:0000`.

#### Output

Output is the size of shortest list of CIDR masks that cover all IP addresses except those listed in input.

The CIDR should be formatted as any representation of an IP followed by / and amount of bits in the prefix (in decimal). For example, `0243:F6A8:885A:308D:3131:98A2:0000:0000/112`.
You are not required to shorten the IP at all, as in above example. However `243:f6a8:885a:308d:3131:98a2::/112` is not invalid.

### Example

```go
NumberOfMasks([]string{"0000:0000:0000:0000:0000:0000:0000:0000"}) // 128
NumberOfMasks([]string{"0000:0000:0000:0000:0000:0000:0000:0000", "0000:0000:0000:0000:0000:0000:0000:0001"}) // 127
```

### Run tests with benchmarks

```
go test -bench .
```
