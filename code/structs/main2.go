// With a slice (example of reference type different that with a struct example of value types  (basics + structs) we don't have to pass address to allow modify value of elements of the slice. The reason is because although reference types are passed to a function => go copy reference type this reference types inside pointer to underlying array keeps pointing to same array!!

package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "there", "How", "Are", "You"}
	
	updateSlice(mySlice)
	
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

/*
Output:
E:\DERMS\Curses\5_Alvaro_Basic_Go\GoCasts\code\structs>go run main2.go
[Bye there How Are You]
*/
