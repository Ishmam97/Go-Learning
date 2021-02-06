//recursive search
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 99, 24, 56, 49, 36, 69}
	startIdx := 0
	target := 24
	// return the index of target item. Returns -1 if not found
	fmt.Println(recursiveSearch(arr, startIdx, target))
}
func recursiveSearch(arr []int, startIdx, target int) int {
	if startIdx == len(arr) {
		return -1
	} else if arr[startIdx] == target {
		return startIdx
	} else {
		return recursiveSearch(arr, (startIdx + 1), target)
	}

}
