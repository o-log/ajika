package wraperr

import (
	"fmt"
	"runtime"
)

// Appends to the end of the error new line with caller file, line and function name.
// Must be used when proxying errors from underlying functions to enrich error with trace of functions it goes through.
// Uses go "errors wrapping": https://go.dev/blog/go1.13-errors
// Resulting error objects can be correctly unwraped.
func Wrap(err error) error {
	var functionName string
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			functionName = fn.Name()
		}
	}

	return fmt.Errorf("%w\n%v:%v %v", err, file, line, functionName)
}

// Assembles error message with caller file, line and function name.
func Err(comment string) error {
	var functionName string
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			functionName = fn.Name()
		}
	}

	return fmt.Errorf("%v:%v %v %v", file, line, functionName, comment)
}
