// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	switch {
	case true:
		fmt.Println(`es true`)
	case false:
		fmt.Println(`es false`)
	default:
		fmt.Println(`?`)
	}
}

// STOPMAIN OMIT
