package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// Type Conversion type(value) => int(y) = 10.0 => 10

	fmt.Println(x + int(y))

	// ASCII
	num1 := 106
	// str1 := strconv.Itoa(num1)
	str1 := rune(num1)
	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", str1, str1)
	fmt.Printf("Character corresponding to Ascii Code %d = %c\n", num1, str1)
}
