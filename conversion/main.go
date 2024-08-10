package main

import (
	"fmt"
)

func main() {
	var name = "edbert"
	var charName byte = name[0]
	fmt.Printf("char = %c\n", charName)

	var unsigned32 uint32 = 10
	var unsigned8 uint8 = uint8(unsigned32)

	var charString string = string(unsigned8)

	fmt.Printf("char = %s\n", charString)
}
