package test

import (
	"fmt"
	"runtime"
)

func Version() string {
	fmt.Println(runtime.Version())
	return runtime.Version()
}
