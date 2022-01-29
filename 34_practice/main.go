// 1-) Underlying Type 'int' olacak şekilde kendi veri tipinizi oluşturunuz veri tipini ve değerini 10 ile yazdırınız

// package main

// import "fmt"

// type myVeri int

// func main() {
// 	var x myVeri
// x = 10

// 	fmt.Printf("%T %#v\n", x, x)
// }

// 2-) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.

// package main

// import "fmt"

// func main() {
// 	x := 10
// 	fmt.Println(x)
// 	y := &x
// 	*y = 20
// 	fmt.Println(x)
// }

// 3-) Underlying Type struct olan Rectrangle type oluşturunuz. display,area, circumference metdolarını yazınız.

// package main

// import "fmt"

// type Rectrangle struct {
// 	a, b int
// }

// func (r Rectrangle) display() {
// 	fmt.Println("Kenar uzunlukları: ", r.a, " ve ", r.b, "olan dikdörtgen")
// }
// func (r Rectrangle) area() int {
// 	return r.a * r.b
// }
// func (r Rectrangle) circumference() int {
// 	return 2 * (r.a * r.b)
// }
// func main() {
// 	r1 := Rectrangle{5, 7}
// 	r1.display()
// 	fmt.Println("Alanı: ", r1.area())
// 	fmt.Println("Çevresi", r1.circumference())
// }

// 4-) 4 ader user ı struct yapıyısyl farklı şekilde tanımlayınız. name,age,int for döngüsü kullanıcları gösteriniz
package main

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {

	family1 := family{
		name:      "Arin",
		age:       5,
		isMarried: false,
	}

	family2 := family{
		name: "Elis",
		age:  3,
	}

	family3 := family{
		"Gurcan",
		40,
		true,
	}

	var family4 family
	family4.name = "Gamze"
	family4.age = 40
	family4.isMarried = true

	return []family{family1, family2, family3, family4}
}

func main() {
	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
	}
}
