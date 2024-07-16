package main

import "fmt"

// Pada bagian ini kita akan menjadikan sebuah function itu merupakan value dari sebuah variable
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := getGoodBye // artinya, contoh merupakan variable yang mengacu pada function getGoodBye()

	fmt.Println(contoh("Miftahul"))
}