package main

// import (
// 	"errors"
// )

// type customer struct {
// 	id      int
// 	balance float64
// }

// type transactionType string

// const (
// 	transactionDeposit    transactionType = "deposit"
// 	transactionWithdrawal transactionType = "withdrawal"
// )

// type transaction struct {
// 	customerID      int
// 	amount          float64
// 	transactionType transactionType
// }

// // Don't touch above this line

// func updateBalance(cust *customer, tx transaction) error {
// 	if tx.transactionType == transactionWithdrawal {
// 		newBalance := cust.balance - tx.amount
// 		if newBalance < 0 {
// 			return errors.New("insufficient funds")
// 		}
// 		cust.balance = newBalance
// 		return nil
// 	} else if tx.transactionType == transactionDeposit {
// 		cust.balance += tx.amount
// 		return nil
// 	}
// 	return errors.New("unknown transaction type")
// }
