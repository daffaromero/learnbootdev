package main

func adder() func(int) int {
	sum := 0
	return func(add int) int {
		sum += add
		return sum
	}
}
