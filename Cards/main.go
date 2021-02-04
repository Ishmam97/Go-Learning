package main

func main() {
	cards := deck{}

	cards = append(cards, "Six of spades")
	cards.print()

}
func newCard() string {
	return "5 of diamonds"
}
