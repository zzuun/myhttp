# myhttp

This tool takes different addresses and print hash of their responses by making http calls.


## How to run

```go run main.go http://www.adjust.com http://google.com```

This tool also supports concurrency. You can pass how many parallel request should be made by default it's 10

### Example
```go run main.go http://www.adjust.com http://google.com --parallel 5```





