package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name
}

func getSum(num1 int, num2 int) (int, int, string){
	return num1 +num2, num1, "CC"
}

func main(){
	fmt.Println(greeting("VinhHoang97"))
	fmt.Println(getSum(2,3))
}


