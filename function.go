package main

import "fmt"

func sayHello() {
	fmt.Println("Halo Bang")
}
// FUNCTION DENGAN PARAMETER
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

//FUNCTION RETURN VALUE VERSI 1
func getHello(name string)string{
	return "Hello " + name //Menggunakan kata kunci "return"
}
//FUNCTION RETURN VALUE VERSI 2
func getHello1(name string)string{
	result := "Hai " + name
	return result
}

func main() {
	sayHello()
	sayHelloTo("Miftahul","Falah") //wajib mengirimkan semua parameter yang diminta 
	result:=getHello("Miftahul") //ini versi yang disimpan dalam variabel 
	fmt.Println(result)
	getHello1("Falah") // ini yang variabelnya disimpan di function
	fmt.Println(getHello1("Mario")) //ini versi kalo kita gamau nyimpan di dalam variabel
}