package main // Organize the code by pack

import "fmt"

/*
- `int`: 0
- `float`: 0.0
- `bool`: false
- `string`: "" (empty string)
- `pointer`: nil
- `struct`: all fields are set to the zero value for their respective types
*/
func main() {
	var x int //0
	var arr [5]int // [0,0,0,0,0]
	var arr [5]string // ["", "", "", "", ""]
	var slice []int
	fmt.Println(slice == nil) // => true
}

