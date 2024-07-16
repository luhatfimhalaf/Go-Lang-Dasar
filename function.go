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

// FUNCTION MULTIPLE VALUE 
func getFullName1() (string,string){
	return "Miftahul", "Falah"
}

// CARA MENGHIRAUKAN RETURN VALUE
func getFullName2()(string,string){
	return "Luhatfim","Halaf"
}

// KODINGAN NAMED RETURN VALUES
// membuat variable secara langsung di tipe data return function nya
func getCompleteName()(firstName3,middleName3,lastName3 string){ //ini kita bisa singkat seperti ini deklarasi variabelnya (kalau tipe datanya sama semua)
	firstName3 = "Miftahul" //sehingga disini tidak perlu dideklarasikan ulang 
	middleName3 = "Falah"
	lastName3 = "Luhatfim"

	return firstName3,middleName3,lastName3
}

func main() {
	sayHello()
	sayHelloTo("Miftahul","Falah") //wajib mengirimkan semua parameter yang diminta 
	result:=getHello("Miftahul") //ini versi yang disimpan dalam variabel 
	fmt.Println(result)
	getHello1("Falah") // ini yang variabelnya disimpan di function
	fmt.Println(getHello1("Mario")) //ini versi kalo kita gamau nyimpan di dalam variabel
	firstName, lastName := getFullName1() //eksekusi function multiple value
	fmt.Println(firstName,lastName)

	firstName1,_:= getFullName2() //eksekusi jika kita ingin menghiraukan (tidak memakai salah satu return) yaitu dengan cara kita ganti dengan _ (garis bawah)
	fmt.Println(firstName1)

	firstName3,middleName3,lastName3 := getCompleteName()
	fmt.Println(firstName3,middleName3,lastName3)
}