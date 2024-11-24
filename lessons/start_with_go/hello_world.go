// You can edit this code!
// Click here and start typing.
// Double forward slash for comment
/*
For larger comment
*/
package main // Organize the code by pack

import "fmt"

func main() {
	fmt.Println("Hello, Gophers")

	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
	// To find the type of the value we can use %T
	// And using %v as the default type
	// Different from python, we need to add \n in the print in order to get the new lin
	fmt.Printf("%v type is %T\n", name, name)
	fmt.Printf("%v type is %T", age, age)
}

