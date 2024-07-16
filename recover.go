package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("End App")
	message := recover() //Nahh yang bener disini naruhnya 
	fmt.Println("Terjadi Panic", message) 

}

func runApp(error bool){
	defer endApp()
	if error{
		panic("ERROR")
	}

	// message := recover() //Kalau kita masang recover nya disini itu salah akrena pasti defer duluan yang akan dieksekusi
	// fmt.Println("Terjadi Panic", message) 
}

func main()  {
	runApp(true)
}