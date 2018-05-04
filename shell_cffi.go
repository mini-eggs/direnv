package main

import (
	"errors"
	"fmt"
)

type cffi int

var CFFI cffi

func (x cffi) Hook() (string, error) {
	return "", errors.New("this feature is not supported.")
}

func (x cffi) Export(e ShellExport) (out string) {
	// for key, value := range e {
	// 	if value == nil {
	// 		// unset
	// 	} else {
	// 		// set
	// 	}
	// }
	fmt.Printf("\n%+v\n", e)
	// out = "testing"
	return out
}
