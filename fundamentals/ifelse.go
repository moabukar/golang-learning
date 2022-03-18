package main

import "fmt"

func main() {

	age := 13

	if age >= 18 {

		fmt.Println("You can ride alone")

	} else if age > 14 && age < 18 {
		fmt.Println("you can ride with a parent")
	} else {
		fmt.Println("You cannot ride")
	}
	// else {
	// 	fmt.Println("You cannot ride")
	// 	fmt.Printf("Wait %d years", 18-age)
	// }

}
