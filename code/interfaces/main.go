package main

import "fmt"

// We use interfaces to define function signatures for "VERY custom logic" so must use a "common type"
// Define a new type ...
type bot interface { //  ... that will be implemented by (will be able to subst to) both eb & sb
     // Inside interfaz we must write all the functions (& whith input parameters types and return value types) any types must implement in order to be considered as matching this interface type  <=> in order to implement this "inteface type"
	getGreeting() string //  ... and will enable to call printGreeting function that has tha bot type
	//  ... in his signature
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//func printGreeting(eb englishBot) {
//func printGreeting(sb spanishBot) {
func printGreeting(b bot) { // VERY common logic => will delegate in interface
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // VERY custom logic  => to interface
	// VERY custom logic for generating an english greeting => Not reusable their inside code => to interface
	return "Hi there!"
}

func (spanishBot) getGreeting() string { // VERY custom logic  => to interface
	return "Hola!"
}

/*
Output:
---------

Hi there!
Hola!



*/
