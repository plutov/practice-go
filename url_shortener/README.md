### url_shortener

This function can be used to create a URL shortener service. You're required to implement two functions: `Shorten` to convert a long URL into a short, unique code, and `Expand` to retrieve the original URL from the short code.It provides a more scalable and flexible solution

### Shorten URL
```
POST /shorten
curl -X POST -d "url=https://www.example.com" http://localhost:8080/shorten
Shortened URL: zbTYjcoL

```

### Expand URL
```
GET /expand/{short_code}
curl http://localhost:8080/expand/zbTYjcoL
<a href="https://www.example.com">Found</a>.
```

### Run tests with benchmarks

This'll run the tests and give you some fancy benchmarks.

```
go test -bench .
```
