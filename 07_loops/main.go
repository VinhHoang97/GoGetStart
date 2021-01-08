package main

import "fmt"

func main() {
	//Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	for i := 1; i <= 1000; i++ {
		if i%20 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%4 == 0{
			fmt.Println("Fiz")
		}else if i%5 == 0{
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
