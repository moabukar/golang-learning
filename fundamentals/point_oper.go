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

func changeValue(str *string) {
	*str = "changed"
}

func chnageValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"
	fmt.Println(toChange)
	chnageValue2(toChange)
	fmt.Println(toChange)
}
