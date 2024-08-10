package main

import (
	"fmt"
)

func main() {
	type noKtp string
	type age uint8

	var noKtpEdbert noKtp = "1234567890"
	var ageEdbert age = 20

	fmt.Printf("noKtp = %s && age= %d\n", noKtpEdbert, ageEdbert)
}
