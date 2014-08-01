trace
======

Trace package adds caller file name and line number to 
logs and errors.

### Installation
    go get github.com/graarh/golang/trace

### Usage
``` golang
trace.Error("error message")
trace.Log("log something", something)
trace.Logf("%s, %s", "log something", "another way")
```
