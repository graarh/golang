// Package trace adds file and line number to errors and logs
package trace

import (
	"errors"
	"fmt"
	"log"
)

//Error creates new error with file name and line of caller
func Error(v ...interface{}) error {
	v = append(v, fileLine())
	return errors.New(fmt.Sprint(v...))
}

//Errorf is like error but with template
func Errorf(format string, a ...interface{}) error {
	text := fmt.Sprintf(format, a...)
	return errors.New(fmt.Sprint(text, fileLine()))
}

//Log calls log.Println for given arguments and adds file name and line of caller
//Another way to do so is the Lshortfile log flag
func Log(v ...interface{}) {
	v = append(v, fileLine())
	log.Println(v...)
}

//Logf is like Log but with file and line
func Logf(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	log.Print(fmt.Sprint(text, fileLine()))
}

//LogBacktrace is like log.Print but with full callstack backtrace
func LogBacktrace(v ...interface{}) {
	v = append(v, backtrace(2))
	log.Println(v...)
}

//LogBacktracef is like LogBacktrace but with template
func LogBacktracef(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	log.Print(text, backtrace(2))
}
