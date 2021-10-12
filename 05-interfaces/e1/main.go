package main

import "fmt"

type person struct {
	name    string
	address string
	city    string
	state   string
	zip     string
}

func (p person) String() string {
	return fmt.Sprintf("%s | %s, %s, %s %s", p.name, p.address, p.city, p.state, p.zip)
}

type people []person

func (peeps people) String() string {
	var result string
	for _, p := range peeps {
		result += fmt.Sprintf("%s\n", p)
	}
	return result
}

func main() {
	peeps := people{
		{name: "John", address: "123 Main St", city: "Jamestown", state: "NY", zip: "14701"},
		{name: "Jane", address: "234 Fleet St", city: "Columbia", state: "MD", zip: "21150"},
		{name: "Terry", address: "345 Charles Ave", city: "Gergetown", state: "DC", zip: "20007"},
	}

	fmt.Println(peeps)
}
