package main

import "fmt"

func main() {

	//Short Declarative Variable
	person := map[string]string{
		"name":    "Miftahul",
		"address": "Jepara",
	}

	// //Manually Declarative Variable
	// var orang map[string]string = map[string]string{}
	// person["nama"] = "Falah"
	// person["alamat"] = "Kudus"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["salah"])
	
	// fmt.Println(orang)
	// fmt.Println(orang["nama"])
	// fmt.Println(orang["alamat"])

	// FUNCTION MAP
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Miftahul Falah"
	book["wrong"] = "Salahh Pakk"

	fmt.Println(book)
	delete(book,"wrong")

	fmt.Println(book)
}