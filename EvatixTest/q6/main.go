package main

import "fmt"

func main() {
	fmt.Println(cashier(523))
}

func cashier(amount int) *map[int]int {
	change := map[int]int{
		1:    0,
		2:    0,
		5:    0,
		10:   0,
		20:   0,
		50:   0,
		100:  0,
		200:  0,
		500:  0,
		1000: 0,
	}
	num := []int{1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	for _, val := range num {
		if amount < val {
			continue
		} else {
			diff := amount / val
			change[val] = diff
			amount = amount - (val * diff)
		}
	}
	return &change
}
