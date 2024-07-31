### url_shortener

This function can be used to create a URL shortener service. You're required to implement two functions: `Shorten` to convert a long URL into a short, unique code, and `Expand` to retrieve the original URL from the short code.

### Example
```
Enter a URL to shorten or a short code to expand:
https://www.google.com/

Alright, here's your shortened URL: 59e3b40e

Enter a URL to shorten or a short code to expand:
59e3b40e
Here's the original URL: https://www.google.com/
```
### A heads up...
The current implementation uses an in-memory map to store URL mappings, it forgets everything when you close it. All those shortened URLs? Poof! Gone. If you require persistence across runs, consider using a file-based storage solution. For example, you might want to look into using a JSON file to persist URL mappings between program runs.

### Run tests with benchmarks

This'll run the tests and give you some fancy benchmarks.

```
go test -bench .
```
