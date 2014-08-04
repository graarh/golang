package trace

import (
	"fmt"
	"runtime"
)

func fileLine() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf(" <%s:%v>", file, line)
}

func backtrace(skip int) string {
	var stack string

	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		Func := runtime.FuncForPC(pc)
		stack += fmt.Sprintf("\n	at %s(%s:%v)", Func.Name(), file, line)
		if Func.Name() == "main.main" {
			break
		}
	}

	return stack
}
