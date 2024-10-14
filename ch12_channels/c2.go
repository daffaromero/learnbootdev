package main

// import "time"

// func processMessages(messages []string) []string {
// 	ch := make(chan string)
// 	var output []string

// 	for _, msg := range messages {
// 		go func(s string) {
// 			processed := process(s)
// 			ch <- processed
// 		}(msg)
// 	}

// 	for i := 0; i < len(messages); i++ {
// 		msg := <-ch
// 		output = append(output, msg)
// 	}

// 	return output
// }

// // don't touch below this line

// func process(message string) string {
// 	time.Sleep(1 * time.Second)
// 	return message + "-processed"
// }
