package main

import "fmt"

func main() {
	var x int = 25
	var y int = 85
	var sayi1 int
	fmt.Print("x ve y sayilarini sirasiyla giriniz: ")
	sayi1 = ahmet(x, y)
	fmt.Printf("%d", sayi1)
}

func ahmet(a, b int) int {
	var x, y int

	for x = a; x < y; x++ {
		for y = b; y < x; y++ {
			if x%y == 0 {
				break
			}
		}
		if y == x {
			fmt.Printf("%d bir asal sayıdır\n", y)
		}
	}
	return -1
}
