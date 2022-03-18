package main

import "fmt"

// func main() {
// 	var arr [5]int
// 	arr[0] = 100
// 	arr[4] = 200

// 	fmt.Println(arr[0])
// }

// func main() {
// 	arr := [3]string{"test", "hello", "test3"}

// 	fmt.Printf("these are the arrays mate %q", arr)

// }

func main() {
	arr := [3]int{4, 4, 5}
	// fmt.Println(arr)
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)

	// 2D arrays below and you can also create multi-dimensional arrays
	arr2D := [2][3]int{{2, 3, 4}, {3, 4, 5}}
	fmt.Println(arr2D[1][2])
}
