// Package trace adds file and line number to errors and logs
package trace

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

func prepare(v ...interface{}) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s <%s:%v>", fmt.Sprint(v), file, line)
}

//Error creates new error with file name and line of caller
func Error(v ...interface{}) error {
	return errors.New(prepare(v))
}

//Errorf is like error but with template
func Errorf(format string, a ...interface{}) error {
	text := fmt.Sprintf(format, a)
	return errors.New(prepare(text))
}

//Log calls log.Println for given arguments and adds file name and line of caller
func Log(v ...interface{}) {
	log.Println(prepare(v))
}

//Logf is like Log but with template
func Logf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a)
	log.Print(prepare(text))
}
