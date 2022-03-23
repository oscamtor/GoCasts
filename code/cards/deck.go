package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d aqui es un receiver (va antes de nombre_funcion) => podre invocar d.print(). Equivale a un metodo privado de clase deck en C++
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d aqui es un parametro (va despues del nombre_funcion).
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // To convert the []string to a single string we pass it to function Join() that takes it as first parameter & what it does is to "join every value inside of it with a , in between each value" . Ie ["red, "yellow"] --> "red,yellow"
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Randomize the order of cards inside it
func (d deck) shuffle() {
	// r (a random number generator) type is Rand is a source (<=> seed) of random numbers <=> It is an object that will generate random numbers for us

	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	// := rand.NewSource(len(d) - 1)
	// time is an instant.Now() returns a time. time has a method UnixNano() that return that instant as a Unix time: the number of nanoseconds elapsed since January 1, 1970 UTC
	// We will use that time.Now().UnixNano() which is different int64 number as the seed to generate a source object & finally we will use that source objes as the basis for our new random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Podríamos pero NO necesitamos cada card solo el indice
	//for i, card := range d {
	for i := range d {
		// Bad
		// ---
		// newPosition := rand.Intn(len(d) - 1)
		// because the random will be the same even with a seed variable like the one  that time.Now().UnixNano() give us => the last 4 cards after shuffle will be the same in different ejecutions!!
		// Good
		// ----
		// From doc:
		// func (r *Rand) Intn(n int) int
		// tell us that if we have a value of type Rand, that value can call the Int function passig it an int and get back an int as return
		newPosition := r.Intn(len(d) - 1) // For each card (<=> each index) of the slice, we are going to generate a random number between 0 & the max len of the slice <=> len(d)-1

		d[i], d[newPosition] = d[newPosition], d[i] // swap the current card & the card at cards[randomNumber]
	}
}
