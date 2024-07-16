package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

func main() { 
	blacklist := func(name string) bool{ //Kita membuat function di dalam sebuah variabel
		return name == "Pantekkk"
	}
	registerUser("Falah",blacklist)

	registerUser("Pantekkk", func(name string) bool { //Atau bisa kita buat seperti ini
		return name == "Pantekkk"
	}) 
}