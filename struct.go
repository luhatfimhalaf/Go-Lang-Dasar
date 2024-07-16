package main

import "fmt"

type Customer struct {
	Name, Adress string //Saat kita membuat struct, kita membuatnya dengan huruf besar
	Age          int
}

func main() {
	var eko Customer
	eko.Name = "Miftahul Falah"
	eko.Adress = "Indonesia"
	eko.Age = 30

	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Adress)
	fmt.Println(eko.Age)

	Joko := Customer{
		Name:"Joko",
		Address: "Indonesia",
		Age: 30,
	}
	fmt.Println(Joko)

	budi := Customer {"Budi","Indonesia",30}
	fmt.Println(budi)
}