package main

import (
	"fmt"
)

func main() {
	var names = [5]string{
		"edbert",
		"joko",
		"budi",
		"agus",
		"bambang",
	}

	var slice = names[:2]
	fmt.Println(slice)
}
