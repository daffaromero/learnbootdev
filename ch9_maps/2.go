package main

// import "errors"

// func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
// 	userData, ok := users[name]
// 	if !ok {
// 		return true, errors.New("not found")
// 	} else if !userData.scheduledForDeletion {
// 		return false, nil
// 	}
// 	delete(users, name)
// 	return true, nil
// }

// type user struct {
// 	name                 string
// 	number               int
// 	scheduledForDeletion bool
// }
