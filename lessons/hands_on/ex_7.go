package main

import "fmt"

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	c3 = iota // c0 == 0
	c4
	c5
	c6
)
func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
}
//0 1 2
//0 1 2 3