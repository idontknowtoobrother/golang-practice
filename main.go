package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Address *Address
}

func updateAddress(p *Person, street string, city string) {
	p.Address.Street = street
	p.Address.City = city
}

func main() {

	person := Person{
		Name: "Jakkrit",
		Address: &Address{
			Street: "Maple St",
			City:   "Springfield",
		},
	}

	fmt.Printf("Original address: Street: %s, City: %s\n", person.Address.Street, person.Address.City)
	updateAddress(&person, "Elm St", "Metropolis")
	fmt.Printf("Updated address: Street: %s, City: %s\n", person.Address.Street, person.Address.City)
}
