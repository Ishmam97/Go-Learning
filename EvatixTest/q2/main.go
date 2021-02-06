package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6, 5}
	p := &arr
	fmt.Println(sumOfAllWithinRange(p, 1, 5))
}
func sumOfAllWithinRange(items *[]int, startRange int, endRange int) int {
	var sum int
	for i := startRange; i <= endRange; i++ {
		p := &((*items)[i])
		if p == nil {
			return 0
		} else {
			sum += (*items)[i]
		}
	}

	return sum
}
