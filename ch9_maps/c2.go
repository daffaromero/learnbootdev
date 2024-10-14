package main

// func findSuggestedFriends(username string, friendships map[string][]string) []string {
// 	direct, ok := friendships[username]
// 	if !ok {
// 		return nil
// 	}
// 	suggestedMap := make(map[string]bool)
// 	for _, friend := range direct {
// 		for _, indirect := range friendships[friend] {
// 			if indirect != username && !contains(direct, indirect) {
// 				suggestedMap[indirect] = true
// 			}
// 		}
// 	}

// 	var suggested []string
// 	for friend := range suggestedMap {
// 		suggested = append(suggested, friend)
// 	}
// 	return suggested
// }

// func contains(slice []string, item string) bool {
// 	for _, v := range slice {
// 		if v == item {
// 			return true
// 		}
// 	}
// 	return false
// }
