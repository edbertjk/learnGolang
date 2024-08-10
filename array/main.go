package main

import (
	"fmt"
)

func main() {
	var names = [3]string{
		"edbert",
		"joko",
		"budi",
	}

	var age = map[string]uint8{
		"edbert": 15,
		"joko":   16,
		"budi":   17,
	}

	for i := 0; i < len(names); i++ {
		fmt.Printf("name %d = %s (%d Tahun)\n", i, names[i], age[names[i]])
	}
}
