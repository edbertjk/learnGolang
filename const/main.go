package main

import (
	"fmt"
)

func constVar() {
	const (
		a        = 10
		b string = "hello"
		c bool   = false
	)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %s\n", b)
	fmt.Printf("c = %t\n", c)
}
