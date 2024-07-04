package main

import "fmt"

func main() {
	var a int = 30
	var b int = 30

	var c int = a + b
	fmt.Println(c)

//Augmented Assigments
var i = 10
i += 10
fmt.Println(i)

i += 15
fmt.Println(i)

//Unary Operator

var j = 1
j ++ // j = j+1
fmt.Println(j)

j ++ // j = j+1
fmt.Println(j)

j -- //menurunkan
fmt.Println(j)
j --
fmt.Println(j)
}