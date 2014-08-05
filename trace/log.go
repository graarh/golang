// Package trace adds file and line number to errors and logs
package trace

import (
	"errors"
	"fmt"
	"log"
)

//Error creates new error with full backtrace
func Error(v ...interface{}) error {
	v = append(v, backtrace(2))
	return errors.New(fmt.Sprint(v...))
}

//Errorf is like error but with template
func Errorf(format string, a ...interface{}) error {
	text := fmt.Sprintf(format, a...)
	return errors.New(fmt.Sprint(text, backtrace(2)))
}

//Log is like log.Print but with full callstack backtrace
func Log(v ...interface{}) {
	v = append(v, backtrace(2))
	log.Println(v...)
}

//Logf is like LogBacktrace but with template
func Logf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	log.Print(text, backtrace(2))
}
