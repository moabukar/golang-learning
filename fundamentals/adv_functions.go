package main

import "fmt"

// func test() {
// 	fmt.Println("This is the test function.")
// }

// func main() {
// 	x := test
// 	x()
// }

// part 2

// func main() {
// 	test := func() {
// 		fmt.Println("Your test function")
// 	}
// 	test()
// }

// func main() {
// 	test := func(x string) {
// 		fmt.Println(x)
// 	}
// 	test("Your test function")
// }

// func main() {
// 	test := func(x, y int) int {
// 		return x + y
// 	}(7, 4)
// 	fmt.Println(test)
// }

// func main() {
// 	test := func(x int) int {
// 		return x * -2
// 	}
// 	test2(test)
// }

// func test2(myFunc func(int) int) {
// 	fmt.Println(myFunc(7))
// }

// calling functions immediately

// func main() {
// 	func() {
// 		fmt.Println("test")
// 	}()
// }

func main() {
	returnFunc("hello")()
}

// function closure
func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}
