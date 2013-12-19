// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	str := `Hola, привет!`
	for i, r := range str {
		fmt.Printf("str(%d)='%c'\n", i, r)
	}
}

// STOPMAIN OMIT
