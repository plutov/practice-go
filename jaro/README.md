### Jaro distance

The Jaro distance is a measure of similarity between two strings. The higher the Jaro distance for two strings is, the more similar the strings are. The score is normalized such that 0 equates to no similarity and 1 is an exact match.

The Jaro score of 2 given strings is:

![Jaro Formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/ba49d2ef630a599848c412d62e62647edbaeb306)

Where:
 - `m` is the number of matching characters
 - `t` is half the number of transpositions

Write a function `Distance(word1 string, word2 string) float64` that will find Jaro distance.

Sample Jaro distance:
```
jones, johnson, 0.790476
```

### Run tests with benchmarks

```
go test -bench .
```
