package main

import "fmt"

func main() {
	ids:= []int{33,76,54,23,11,2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n",  id)
	}

	sum:= 0
	for _,id:=range ids{
		sum+=id
	}
	fmt.Println("sum", sum)

	emails := map[string]string{"bob": "bobmail", "sharon": "sharonmail"}
	for key,value := range emails {
		fmt.Printf("%s: %s\n",key, value)
	}

	for k:= range emails{
		fmt.Println("Name:" +k)
	}
}
