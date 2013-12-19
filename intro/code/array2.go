// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	x := [...]int{1, 2, 3, 4, 5}
	fmt.Println(x[:3], x[2:], x[2:4])
}

// STOPMAIN OMIT
