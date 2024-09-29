package main

// import "fmt"

// func (e email) cost() int {
// 	if e.isSubscribed {
// 		return 2 * len(e.body)
// 	}
// 	return 5 * len(e.body)
// }

// func (e email) format() string {
// 	return fmt.Sprintf("'%s' | %s", e.body, func() string {
// 		if e.isSubscribed {
// 			return "Subscribed"
// 		}
// 		return "Not Subscribed"
// 	}())
// }

// type expense interface {
// 	cost() int
// }

// type formatter interface {
// 	format() string
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// }
