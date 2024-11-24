package main

import (
	"fmt"
)

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it 
		// in a variable that is declared to hold a VALUE of float64
		z = m 
		fmt.Printf("%v of type %T \n", z, z)
	*/

	/*
		// this does work
		z = float64(m)
		fmt.Printf("%v of type %T \n", z, z)
	*/
	// if x := f(); x < y {
	// 	return x
	// } else if x > z {
	// 	return z
	// } else {
	// 	return y
	// }

	// this is a struct call Address
	type Address struct {
		Street  string
		City    string
		Country string
	}
	// this is a struct call Person 
	type Person struct {
		Name    string
		Age     int
		Address Address
	}

	p := Person{
		Name: "John",
		Age:  25,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	fmt.Println("Street:", p.Address.Street)
	fmt.Println("City:", p.Address.City)
	fmt.Println("Country:", p.Address.Country)

	/*
	Street: 123 Main St
	City: New York
	Country: USA
	*/

}
