// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
type Vector [3]int32

func (v Vector) String() string {
	return fmt.Sprintf(`v=(%d, %d, %d)`, v[0], v[1], v[2])
}

func main() {
	x := Vector{1, 2, 3}
	fmt.Println(x)
}

// STOPMAIN OMIT
