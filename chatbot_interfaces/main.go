package main

import "fmt"

type bot interface {
	getGreeting() string
}
type engBot struct{}
type espBot struct{}

func main() {
	eb := engBot{}
	sb := espBot{}
	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (engBot) getGreeting() string {
	//complex logic here
	return "hi there"
}

func (espBot) getGreeting() string {
	//complex logic here
	return "hola Ninoo"
}
