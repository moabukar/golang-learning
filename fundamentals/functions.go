package main

import "fmt"

// func add(x, y int) {
// 	fmt.Println(x + y)
// }

// func add(x, y int) int {
// 	return x + y
// }

// func add(x, y int) (int, int) {
// 	return x + y, x - y
// }

func add(x, y int) (z1, z2 int) {
	defer fmt.Println("answer done") // defer basically runs after your function exits or when its compelete
	z1 = x + y
	z2 = x - y
	fmt.Println("pre-return")
	return
}

func main() {
	ans1, ans2 := add(30, 15)
	fmt.Println(ans1, ans2)
}
