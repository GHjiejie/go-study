package main

import (
	"fmt"
)

func main() {
	var a *int
	fmt.Println(a) // <nil>
	// fmt.Println(*a) // panic: runtime error: invalid memory address or nil pointer dereference
}
