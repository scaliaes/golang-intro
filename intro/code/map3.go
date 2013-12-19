// +build OMIT

package main

import (
	"fmt"
)

const (
	CHULI     = 1
	KILOCHULI = 1e3
	MEGACHULI = 1e6
)

// STARTMAIN OMIT
func main() {
	molonometer := map[string]int32{`Cubert`: 40 * MEGACHULI} // HL
	if v, ok := molonometer[`Zoidberg`]; ok {
		fmt.Println(v)
	} else {
		fmt.Println(`Y U NO COOL`)
	}
}

// STOPMAIN OMIT
