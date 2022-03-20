package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// Output
	// [72 105 32 116 104 101 114 101 33]

	//cards := newDeck()
	//fmt.Println(cards.toString())
	// Output
	//Ace of Spades,Two of Spades,Three of Spades,Four of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs

	//cards := newDeck()
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()
	//Output:
	// 0 Ace of Spades
	// 1 Two of Spades
	// 2 Three of Spades
	// 3 Four of Spades
	// 4 Ace of Diamonds
	// 5 Two of Diamonds
	// 6 Three of Diamonds
	// 7 Four of Diamonds
	// 8 Ace of Hearts
	// 9 Two of Hearts
	// 10 Three of Hearts
	// 11 Four of Hearts
	// 12 Ace of Clubs
	// 13 Two of Clubs
	// 14 Three of Clubs
	// 15 Four of Clubs
}
