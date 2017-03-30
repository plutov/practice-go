### Short Hash

This function can be used to generate short unique hashes, for example in URL shorteners. You're given a `dictionary` of unique characters and `max length` of result hash. Please implement a function `GenerateShortHashes(dictionary string, maxLen int) []string` which generates all possible, unique hashes with minimum length 1 and maximum lenght `maxLen`.

### Examples

```
GenerateShortHashes("ab", 1)
// []string{"a", "b"}

GenerateShortHashes("ab", 2)
// []string{"a", "b", "aa", "bb", "ab", "ba"}

GenerateShortHashes("ab", 3)
// []string{"a", "b", "aa", "bb", "ab", "ba", "aaa", "baa", "aba", "aab", "bbb", "abb", "bab", "bba"}
```

Note: sequence of hashes doesn't matter.

### Run tests with benchmarks

```
go test -bench .
```
