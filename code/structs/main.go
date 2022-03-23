// With a struct (different that with a slice) we need to pass address to allow modify value of properties of the struct 
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo <=> contactInfo (remove contact the field name, it assumes contactInfo is both the variable name & the type) but then you must use "contactInfo: contactInfo{" instead of "contact: contactInfo{" b
	contactInfo // embedded struct
}

func main() {

	// jim : = {"Jim","Party"}
	// or
	// var jim person
	// jim.firstNAme = "Jim"
	// jim.lastName = "Party"
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@gmail.com",
		// 	zipCode: 94000,
		// },
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,  // Finish with comma event the last propertie key:value
		},
	}
	
	//(B) 1 of 4
	//jimPointer := &jim
	// (B) 2 of 4
	//jimPointer.updateName("jimmy")

	// (C) pointer shortcut (pass a variable in call to a receiver defined as pointer). Is legal and save code!
	jim.updateName("jimmy")
	// fmt.Println(alex)
	// fmt.Println("%+v", alex) // proverty name - value
	jim.print()
}

// (A) Default in Go Pass by value 1 of 2
//func (p Person) updateName(newFirstName string) {    
// (B) 3 of 4
func (pointerToPerson *person) updateName(newFirstName string) { // Read the receiver type as a pointer that points at a person
    //p.firstName = newFirstName						// (A) Default Pass by value 2 of 2
	// (4 of 4)
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}


/*
Output:

E:\DERMS\Curses\5_Alvaro_Basic_Go\GoCasts\code\structs>go run main.go
{firstName:jimmy lastName:Party contactInfo:{email:jim@gmail.com zipCode:94000}}

*/




