package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6, 5, 0, 8}
	findMissing(arr)
}

func findMissing(arr []int) {
	sum := 0
	idx := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			idx = i
		} else {
			sum = sum + arr[i]
		}
	}
	total := (len(arr) + 1) * len(arr) / 2
	fmt.Println("missing number ", (total - sum), "idx ", idx)
}
