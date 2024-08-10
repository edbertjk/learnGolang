package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	var c int = 30

	fmt.Printf("Kali: %d\n", a*b*c)
	fmt.Printf("Tambah: %d\n", a+b+c)
	fmt.Printf("Kurang: %d\n", a-b-c)
	fmt.Printf("Bagi: %d\n", a/b/c)
	fmt.Printf("Sisa Bagi: %d\n", a%b%c)
}
