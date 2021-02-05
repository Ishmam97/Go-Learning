package main

import "fmt"

func main() {
	colors := map[string]string{ //keys type string vals type string
		"red":   "#ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}
	fmt.Printf("%+v", colors)
}
