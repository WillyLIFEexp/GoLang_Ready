package main
// This kind of import works too
// import fmt
// import math/rand
// when importikng specfic library from go, we use / instead of . like we did in python
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}