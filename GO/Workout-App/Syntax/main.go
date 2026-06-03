package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() { // begin main insertion point

	myPerson := Person{
		Name: "John Person; ",
		Age:  25,
		Address: Address{
			Street: " 123 Nowhere ville; ",
			City:   " Somewhere Town",
		},
	}

	fmt.Printf("prior name: %+v\n", myPerson)
	modifyPersonName(&myPerson)
	fmt.Println("new name: ", myPerson.Name)

} // end main insert point

func modifyPersonName(myPerson *Person) {
	myPerson.Name = "Jacob Pine"
}
