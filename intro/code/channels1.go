// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
	c := make(chan int)
	go func() { c <- 1 }()
	go func() { c <- 2 }()
	x, y := <-c, <-c
	fmt.Println(x, y)
}

// STOPMAIN OMIT
