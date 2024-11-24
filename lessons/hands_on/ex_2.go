package main

import "fmt"

// We need  { } to close up the function in GO
func add(x int, y int) int {
	return x + y
}
/*
	Return
	55
*/

// The type after the ) is the return type of the function
func add_2(x, y int, n string) string {
	fmt.Println(n)
	return n
}
/*
	Return
	HELLO WORLD
	HELLO WORLD
*/

func main(){
	// The following expression are the same]
	// x int, y int
	x, y int

}

func swap(x, y string) (string, string) {
	return y, x
}

// Since we putting the return variable after the function, we don't need to add
// the name at the return. 
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}