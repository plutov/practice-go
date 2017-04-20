### Web Scraping

Create a function that finds the time from this [http://tycho.usno.navy.mil/cgi-bin/timer.pl](http://tycho.usno.navy.mil/cgi-bin/timer.pl) and then prints it by extracting the time by timezone code.

### Examples

```
//Apr. 19, 12:59:44 UTC
GetTime("UTC")
```

### Run tests with benchmarks

```
go test -bench .
```
