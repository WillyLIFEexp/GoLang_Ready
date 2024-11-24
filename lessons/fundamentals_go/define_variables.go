package main // Need to define the package we are using

import "fmt" // Import the necessary libraies

func main() {
	/*
	- Different from Python, when declaring variable we need to give a specific type
	- := will let GoLang to automatically identify the variable type
	- := can only defind inside the function, var i <type> can be define inside or outside the function
	*/
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	// if the value been declare then must be use
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	// We can still declare the new value by certain type
	/*
		var g int
		fmt.Println(g)
		g = 43
		fmt.Println(g)

		// declare a variable to hold a VALUE of a certain TYPE
		// initializing a variable
		var h int = 44
		fmt.Println(h)
	*/
}
