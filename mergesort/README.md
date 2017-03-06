### Merge Sort

The [merge sort](https://en.wikipedia.org/wiki/Merge_sort) is a recursive sort of order `n*log(n)`. The basic idea is to split the collection into smaller groups by halving it until the groups only have one element or no elements. Then merge the groups back together so that their elements are in order. This is how the algorithm gets its `divide and conquer` description.

Please write a function `MergeSort(input []int) []int` to sort integers from the lowest to the highest. 

### Run tests with benchmarks

Tests are only checking if input is sorted in a correct order, they can't determine the sorting algorithm, so we will check it manually.

```
go test -bench .
```
