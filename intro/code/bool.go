// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	x := true
	if x == true {
		fmt.Println(`Verdadero`)
	} else {
		fmt.Println(`Falso`)
	}
}

// STOPMAIN OMIT
