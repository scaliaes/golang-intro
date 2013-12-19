// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	m := make(map[string][]int32)
	m[`hola`] = []int32{1, 2}
	m[`otro`] = []int32{3, 4}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// STOPMAIN OMIT
