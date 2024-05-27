### Build a compression tool

This challenge focuses on creating a simple lossless compression tool from scratch. You'll need to implement an algorithm to achieve compression without relying on external libraries, you can use a known algorithm though (for example Huffman). The goal is to balance compression ratio (how much smaller the compressed file is compared to the original) and compression/decompression speed.

This challenge will help you understand the concepts of data compression, explore different algorithms, and practice efficient data structure design. You'll also gain experience in balancing performance and resource usage.

### Sample data set

This challenge includes a sample `dataset.txt` (3.2MB) file which will be used in benchmarking.

### Functions

- `Encode(s string) string`
- `Decode(s string) string`

### Run tests with benchmarks

```
go test -bench .
```
