// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func pow(x float64) (x2 float64, x3 float64) {
	x2 = x * x
	x3 = x2 * x
	return
}
func main() {
	fmt.Println(pow(2))
}

// STOPMAIN OMIT
