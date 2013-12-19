// +build OMIT

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// STARTMAIN OMIT
func main() {
	var z complex128 = cmplx.Pow(math.E, -1i*math.Pi)
	fmt.Printf("z = %f.\n", z)
}

// STOPMAIN OMIT
