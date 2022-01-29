package main

import "fmt"

func main() {
	/* var (
		name      string = "Ahmet"
		age       int    = 40
		isMarried bool   = false

		weight float32 = 72.5
		height int     = 172
	) */

	// var name, age, isMarried, weight, height = "Ahmet", 40, false, 72.5, 172

	name, age, isMarried, weight, height := "Ahmet", 40, false, 72.5, 172

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)

	// zero value
	// string -->"", numeric --> 0,boolean --> false olur
	var class string
	fmt.Println(class)
}
