trace
======

Trace package adds full stack backtrace to log message or error text.

### Installation
    go get github.com/graarh/golang/trace

### Usage
``` golang
trace.Error("error message")
trace.Errorf("error %s", "message")
trace.Log("something")
trace.Logf("%s %s", "aaa", "bbb")
```

Output example:
```
2014/08/05 11:18:55 aaaa bbbb cccc
	at main.A(/usr/local/go/src/test.go:9)
	at main.B(/usr/local/go/src/test.go:17)
	at main.C(/usr/local/go/src/test.go:21)
	at main.main(/usr/local/go/src/test.go:25)```

Compilable program example:
``` golang
package main

import (
	"github.com/graarh/golang/trace"
	"log"
)

func A() {
	trace.Log("aaaa", "bbbb", "cccc")
	trace.Logf("%s %s", "aaa", "bbb")

	log.Print(trace.Error("qqqq", "wwww"))
	log.Print(trace.Errorf("%s %s", "qqq", "www"))
}

func B() {
	A()
}

func C() {
	B()
}

func main() {
	C()
}
```
