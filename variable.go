package main

import "fmt"

func main(){
	var name string

	name = "Miftahul Falah"
	fmt.Println(name)

	name = "Luhatfim Halaf"
	fmt.Println(name)

	//kita bisa menginisialisasikan langsung variable tanpa harus menyebutkan tipe datanya
	var name1 = "Miftahul Falah"
	fmt.Println(name1)

	//kita bisa menyederhanakan pembuatan variable dengan menggunakan :=
	name2 := "Falah Miftahul" //tapi jangan bikin kaya gini 2 kali di variable yang sama karena ya bakal error 
	fmt.Println(name2)

	var(
		firstName = "Miftahul" //di Golang kalo ada variable yang tidak digunakan itu otomatis akan error
		lastName = "Falah"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}

