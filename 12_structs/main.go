package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func (p Person) greet() string {
	return "Hello, my name is "+ p.firstName+ " "+p.lastName+" and i'm " + strconv.Itoa(p.age)+ " years old"
}

func(p *Person) hasBirthday (){
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m"{
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main(){
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{ "Bob",  "Smith",  "Boston",  "m",  25}
	fmt.Println(person1, person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1, person2)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.getMarried("Meow meow")
	person2.getMarried("Woof woof")
}


