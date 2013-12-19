// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
const (
	CHULI     = 1
	KILOCHULI = 1e3 * CHULI
	MEGACHULI = 1e6 * CHULI
)

func main() {
	molonometer := make(map[string]int32)
	molonometer[`Cubert`] = 40 * MEGACHULI
	fmt.Println(molonometer[`Cubert`], molonometer[`Zoidberg`])
}

// STOPMAIN OMIT
