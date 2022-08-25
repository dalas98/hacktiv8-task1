package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	person1 := Human{Name: "Edi", Age: 18, Gender: "Men"}
	person2 := Human{Name: "Fahmi", Age: 22, Gender: "Men"}
	person3 := Human{Name: "Giva", Age: 20, Gender: "Men"}
	person4 := Human{Name: "Clara", Age: 19, Gender: "Women"}
	person5 := Human{Name: "Bayu", Age: 19, Gender: "Men"}
	person6 := Human{Name: "Eka", Age: 20, Gender: "Men"}
	person7 := Human{Name: "Talia", Age: 21, Gender: "Women"}
	person8 := Human{Name: "Irfan", Age: 22, Gender: "Men"}
	person9 := Human{Name: "Fiqri", Age: 23, Gender: "Men"}
	person10 := Human{Name: "Ricky", Age: 22, Gender: "Men"}

	printHuman(&person1, &person2, &person3, &person4, &person5, &person6, &person7, &person8, &person9, &person10)

}

var printHuman = func(human ...*Human) {
	for _, v := range human {
		fmt.Printf("Name: %s, Age: %d, Gender: %s\n", v.Name, v.Age, v.Gender)
	}
}
