package main

import "fmt"

func main() {

	//var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Difference with struct. With maps we don't need to know all the keys at compile time.
	colors["yellow"] = "#FFFF00" // google search: yellow hex code

	delete(colors, "yellow")

	// (A)
	//fmt.Println(colors)

	// (B)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/*
E:\DERMS\Curses\5_Alvaro_Basic_Go\GoCasts\code\map>go run main.go
Hex code for red is #ff0000
Hex code for green is #4bf745
Hex code for white is #ffffff
*/
