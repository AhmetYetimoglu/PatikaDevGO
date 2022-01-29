package main

import "fmt"

func main() {
	var x, y int
	var n int
	for x = 2; x < n; x++ {
		for y = 2; y < x; y++ {
			if x%y == 0 {
				break
			}
		}
		if y == x {
			fmt.Printf("%d bir asal sayıdır\n", y)
		}
	}
}
