package main

import "fmt"

// func main() {
// 	fmt.Println("DATA!")

// 	var x int = 5
// 	var y int = 8
// 	y = 7
// 	fmt.Println(x, y)
// }

// func main() {

// 	var x []int = []int{3, 4, 5}
// 	y := x
// 	y[0] = 100
// 	fmt.Println(x, y)
// }

// func main() {

// 	var x map[string]int = map[string]int{"apple": 1}
// 	y := x
// 	y["y"] = 100
// 	fmt.Println(x, y)
// }

// the above examples show that MAPS and SLICES are MUTABLE data types

// func main() {

// 	var x [2]int = [2]int{4, 5}
// 	y := x
// 	y[1] = 3
// 	fmt.Println(x, y)
// }

// the above example shows a slice type not changing and makes a copy
// MAPS, SLICES and ARRAYS are all MUTABLE types

func main() {
	var x []int = []int{3, 4, 5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}

func changeFirst(slice []int) {
	slice[0] = 1000
}
