### Secret Message

Create a function to decode a secret message, to do it you have to:
 - Sort the characters in the encoded string by the number of times this character appears in it (descending).
 - Now take the sorted string, and drop all the characters after (and including) the `_`. The remaining word is the answer.

### Examples

```
b_bcb_ => b_c => b
```

### Run tests with benchmarks

```
go test -bench .
```
