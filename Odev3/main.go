package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x, y int
	var n int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("n sayısını giriniz: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println()
	n = int(input)
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
