package main

import "fmt"

// func main() {
// 	arr := [3]int{4, 4, 5}
// 	// fmt.Println(arr)
// 	sum := 0

// 	for i := 0; i < len(arr); i++ {
// 		sum += arr[i]
// 	}
// 	fmt.Println(sum)

// 	// 2D arrays below and you can also create multi-dimensional arrays
// 	arr2D := [2][3]int{{2, 3, 4}, {3, 4, 5}}
// 	fmt.Println(arr2D[1][2])
// }

// func main() {

// 	var x [5]int = [5]int{1, 2, 3, 4, 5}

// 	var s []int = x[:3]

// 	fmt.Println(cap(s))
// }

// func main() {

// 	var a []int = []int{5, 6, 7, 8, 9}
// 	a = append(a, 10)

// 	fmt.Println(a)
// }

func main() {

	var a []int = []int{2, 3, 4, 5, 5, 4, 7}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// for i, element := range a {
	// 	fmt.Printf("%d: %d\n", i, element)
	// }
	// for _, element := range a {
	// 	fmt.Printf("%d\n", element)
	// }
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
