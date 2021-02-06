package main

import "fmt"

func main() {
	colors := map[string]string{ //keys type string vals type string
		"red":   "#ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}
	fmt.Printf("%+v\n", colors)
	// delete(colors, "blue")
	// fmt.Printf("%+v\n", colors)
	printMap(colors)
}
func printMap(m map[string]string) {

	for k, v := range m {
		fmt.Println("Hex code for ", k, "is ", v)
	}

}
