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
trace.LogBacktrace("something")
```

Compilable program example and output:
``` golang
package main

import (
	"github.com/graarh/golang/trace"
	"log"
)

func A() {
	trace.Log("aaaa", "bbbb", "cccc")
	trace.Logf("%s %s", "aaa", "bbb")
	trace.LogBacktrace("aa", "bb", "cc")
	trace.LogBacktracef("%s %s", "a", "b")

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

```
2014/08/04 22:39:24 aaaa bbbb cccc  <C:/GoPath/src/test.go:9>
2014/08/04 22:39:24 aaa bbb <C:/GoPath/src/test.go:10>
2014/08/04 22:39:24 aa bb cc
        at main.A(C:/GoPath/src/test.go:11)
        at main.B(C:/GoPath/src/test.go:19)
        at main.C(C:/GoPath/src/test.go:23)
        at main.main(C:/GoPath/src/test.go:27)
2014/08/04 22:39:24 a b
        at main.A(C:/GoPath/src/test.go:12)
        at main.B(C:/GoPath/src/test.go:19)
        at main.C(C:/GoPath/src/test.go:23)
        at main.main(C:/GoPath/src/test.go:27)
2014/08/04 22:39:24 qqqqwwww <C:/GoPath/src/test.go:14>
2014/08/04 22:39:24 qqq www <C:/GoPath/src/test.go:15>
```
