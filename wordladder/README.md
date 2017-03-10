### Word Ladder

Given two words and a dictionary, find the length of the shortest transformation sequence from first word to second word such that:

 - Only one letter can be changed at a time.
 - Each transformed word must exist in the dictionary.
 
Please write a function `WordLadder(from string, to string, dic []string) int` that returns the length of the shortest transformation sequence, or 0 if no such transformation sequence exists.

### Example

```
WordLadder("hot", "dog", []string{"hot", "dog", "cog", "pot", "dot"})
"hot" -> "dot" -> "dog"
3 elements in transformation sequence
```

### Run tests with benchmarks

```
go test -bench .
```
