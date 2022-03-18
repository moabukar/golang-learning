package main

import "fmt"

// func main() {

// 	x := 17
// 	for x > 10 {
// 		fmt.Println(x)
// 		x -= 2
// 	}

// 	// fmt.Println()
// }

// func main() {

// 	for x := 20; x <= 30; x += 2 {
// 		fmt.Println(x)
// 	}
// }

func main() {

	for x := 20; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			break
		}

		// fmt.Println("N ")
	}
}
