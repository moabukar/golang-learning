package main

import "fmt"

// "&" is the reference to the variable - so where the variable is stored in the compute memory

// func main() {
// 	x := 7
// 	y := &x
// 	fmt.Println(x, y)
// 	*y = 8
// 	fmt.Println(x, y)
// }

// Below is an example of changing variables using pointer and dereference operators

func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"
	changeValue2(toChange)
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)
}
