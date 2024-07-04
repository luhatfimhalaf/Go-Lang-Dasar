package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10{
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	// For dengan Statement
	for counter := 1;counter<=11;counter++{
		fmt.Println("Perulangan ke",counter)
	}

	// CARA MANUAL
	names := []string{"Miftahul","Falah","Luhatfim","Halaf"}
	for i:= 0;i<len(names);i++{
		fmt.Println(names[i])
	}
	
	// KODE PROGRAM FOR RANGE
	names1 := []string{"Firstki","Aditya","Fernanda","Utomo"}
	for index, name := range names1{
		fmt.Println("index", index, "=", name)
	}
	names2 := []string{"Firstki","Aditya","Fernanda","Utomo"}
	for _, name := range names2{ //kalau misal kita tidak butuh index nya, kita bisa ganti menjadi _ (garis bawah)
		fmt.Println(name)
	}

}