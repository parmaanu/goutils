package errorutils

import (
	"fmt"
)

// PrintOnErr prints the msg if any of the error is not nil
func PrintOnErr(msg string, errs ...error) bool {
	for _, err := range errs {
		if err != nil && len(msg) > 0 {
			fmt.Println("ERROR, msg", err)
			return true
		}
	}
	return false
}

// PanicOnErr on error panics if error is not nil
func PanicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}
