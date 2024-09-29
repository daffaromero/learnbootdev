package main

import "fmt"

func two() {
	var smsSendingLimit int = 0
	var costPerSMS float64 = 0.0
	var hasPermission bool = false
	var username string = ""

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
