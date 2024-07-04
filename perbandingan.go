package main

import "fmt"

func main() {
	var name1 = "Falah"
	var name2 = "Falah"

	var result1 bool = name1 != name2
	var result2 bool = name1 == name2
	fmt.Println(result1)
	fmt.Println(result2)
	
	//Apakah string bisa mengecek lebih dari atau kurang dari ? bisa, yang di cek adalah urutan alphabet nya 
	
	var alphabet1 string = "a"
	var alphabet2 string = "b"
	
	var result3 bool = alphabet1 > alphabet2
	var result4 bool = alphabet1 < alphabet2
	fmt.Println(result3)
	fmt.Println(result4)

}