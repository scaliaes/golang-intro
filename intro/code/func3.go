// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func Generator(x int) func() int {
	return func() int { return x }
}
func main() {
	f := Generator(3)
	fmt.Println(f())
}

// STOPMAIN OMIT
