// package main

// import "fmt"

// func main() {
// 	deckSize := 20
// 	fmt.Println(deckSize)
// }

// package main

// import "fmt"

// func main() {
// 	var deckSize int
// 	deckSize = 52
// 	fmt.Println(deckSize)
// }

package main

import "fmt"

// var deckSize int

// func main() {
// 	deckSize = 50
// 	fmt.Println(deckSize)
// }

// func main() {
// 	a := []int{5, 4, 3, 2, 1}
// 	a = append(a, 213)

// 	fmt.Println(a)
// }

// using maps below

// func main() {
// 	vertices := make(map[string]int)

// 	vertices["triangle"] = 2
// 	vertices["circle"] = 3
// 	vertices["dodecagon"] = 4
// 	fmt.Println(vertices)
// }

// for loops

// func main() {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(i)
// 	}

// }

// while loops

// func main() {

// 	i := 0

// 	for i < 6 {
// 		fmt.Println(i)
// 		i++
// 	}

// }

//

func main() {

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value")
	}
}
