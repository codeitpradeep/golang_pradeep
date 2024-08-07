package main

import "fmt"

func main() {

	var name = "pradeep"
	age := 27
	fmt.Print("name is %v \n", name)
	fmt.Printf("\nname is %v \n", name)
	fmt.Printf("\nage is %T \n", age)
	fmt.Printf("\nage is %v \n", age)

	var str = fmt.Sprintf("name %v - age %v", name, age)

	fmt.Print(str)
}

/*
name is %v
pradeep
name is pradeep

age is int

age is 27
name pradeep - age 27
*/
