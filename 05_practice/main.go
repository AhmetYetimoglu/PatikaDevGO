package main

import "fmt"

func main() {
	// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
	// 3 farklı yöntem ile oluşturup çıktısını yazınız
	var (
		studentName string = "John Doe"
		grade       int    = 77
		isPassed    bool   = true
	)
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

	// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız
	// studentName, grade, isPassed := "John Doe", 77,true
	// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız.
	// 4-) "Statically Typed" vs "Dynmacilaly Typed" ifadelerini GO ve Python üzerinden gösteriniz.
	// 5-) ":=" vs "=" aradaki farkı gösteriniz.
}
