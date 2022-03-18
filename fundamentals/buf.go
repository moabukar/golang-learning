package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your DOB: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("You typed: %d", 2022-input)
}
