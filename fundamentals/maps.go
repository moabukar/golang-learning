package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apples": 2,
		"pear":   3,
		"orange": 4,
	}

	mp["orange"] = 7
	// mp2 := make(map[string]int)
	mp["test"] = 8
	delete(mp, "apples")
	fmt.Println(mp)
	fmt.Println(mp["orange"])

	// check if key exists in a map

	val, ok := mp["pear"]
	fmt.Println(val, ok)
	fmt.Println(len(mp))
}
