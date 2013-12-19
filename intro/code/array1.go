// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	var x [3]int
	y := [...]float32{3, 2, 4}
	fmt.Println(x, y)
}

// STOPMAIN OMIT
