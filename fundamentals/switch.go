package main

import (
	"fmt"
)

// func main() {

// 	ans := 10

// 	switch ans {
// 	case 1:
// 		fmt.Println("one")

// 	case 2:
// 		fmt.Println("two")

// 	default:
// 		fmt.Println("not 10")
// 	}
// }

func main() {

	ans := 0

	switch {
	case ans > 0:
		fmt.Println("greater than 0")

	case ans < 0:
		fmt.Println("less than 0")

	default:
		fmt.Println("zero")
	}
}
