package main

import "fmt"

func main() {
	fmt.Println(concat("Hello ", 65))
}

func concat(inputString string, inputNumber int) string {
	return fmt.Sprint(inputString, inputNumber)
}
