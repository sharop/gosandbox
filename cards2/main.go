package main

import "fmt"

type contactInfo struct{}
type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	alex.firstName = "Alexander"
	alex.lastName = "Magnus"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)

}
