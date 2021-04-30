package main

import "fmt"

func main() {
	firstName := "John"
	var p *string
	p = &firstName
	fmt.Printf("firstNameのポインター: %d\n", p)

	updateName(firstName)
}

func updateName(name string) {
	name = "David"

	var updated_p *string
	updated_p = &name

	fmt.Printf("updateNameのポインター: %d\n", updated_p)
}
