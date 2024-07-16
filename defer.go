package main

import "fmt"

func logging()  {
	fmt.Println("Selesai memanggil function")
}

func runApplication()  {
	defer logging() // ini dieksekusinya jadinya di akhir
	fmt.Println("Run Application")
}

func main()  {
	runApplication()
}