// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func pow(x float64) (x2, x3 float64, e error) {
	if x > 10 {
		return 0, 0, fmt.Errorf("El valor x=%f es demasiado grande", x)
	}

	return x*x, x*x*x, nil
}

func main() {
	x, y, err := pow(4)
	if err != nil {
		fmt.Printf("ERROR: %v.\n", err)
	} else {
		fmt.Println(x, y)
		x, y = y, x
		fmt.Println(x, y)
	}
}

// STOPMAIN OMIT
