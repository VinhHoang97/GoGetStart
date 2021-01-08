package main

import "fmt"



func main(){

	//var name = "VinhHoang97"
	var age int32 = 24
	const constCool = true

	name:="VinhHoang97"
	var size float32= 1.3

	name, email := "Brad", "email@gmail"

	fmt.Println(name, age, true, size, email)
	fmt.Printf("%T\n", size)
}
