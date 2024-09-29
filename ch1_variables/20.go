package main

import "fmt"

func twenty() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.2f%%", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
}
