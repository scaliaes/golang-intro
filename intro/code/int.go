// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	var x uint32 = 42
	var y uint8 = x
	fmt.Printf("y vale %d.\n", y)
}

// STOPMAIN OMIT
