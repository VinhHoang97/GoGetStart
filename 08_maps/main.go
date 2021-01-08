package main

import "fmt"

func main(){
	email:= make(map[string]string)

	email["NhatRau"] = "nhatrau@gmail.com"
	email["NamRot"] = "namrot@gmail.com"

	fmt.Println(email)
	fmt.Println(email["NhatRau"])
}
