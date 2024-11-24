package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary is %b\n", adams)
	fmt.Printf("42 as hexadeciaml is %x", adams)

	a, b, c := 0, 1, 2
	fmt.Printf("a as binary is %b\n", a)
	fmt.Printf("a as hexadeciaml is %x\n", a)

	fmt.Printf("b as binary is %b\n", b)
	fmt.Printf("b as hexadeciaml is %x\n", b)

	fmt.Printf("c as binary is %b\n", c)
	fmt.Printf("c as hexadeciaml is %x\n", c)
}
