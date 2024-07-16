package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) { //jadi disini si filter itu berupa function
	filteredName := filter(name)
	fmt.Println("Hello ",filteredName)
}

//FUNCTION TYPE DECLARATIONS
// Tujuannya adalah untuk mempersingkat kita dalam penyebutan function parameter
type Filter func(string) string //dan ini untuk type declarationnya
func sayHelloWithFilter1(name1 string, filter Filter){ //nah persingkatnya disini dari type declaration
	filteredName := filter(name1)
	fmt.Println("Mashhookk", filteredName)
}

func spamFilter(name string) string{
	if name == "Anjinggg"{
		return "..."
	}else{
		return name
	}
}
func main() { //cara memanggil function sayHelloWithFilter
	sayHelloWithFilter("Miftahul",spamFilter) //dengan cara langsung seperti ini bisa, dan ingett jangan dipanggil "spamFilter()" karena balikannya nanti berupa string

	filter:= spamFilter //atau kita bisa simpan function nya sebagai values 
	sayHelloWithFilter("Anjinggg",filter)

	sayHelloWithFilter1("Anjinggg",spamFilter)

}