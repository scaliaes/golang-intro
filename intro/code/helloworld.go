// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(`Hello world`)
	}
}

// STOPMAIN OMIT
