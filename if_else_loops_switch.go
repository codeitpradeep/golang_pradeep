package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("if statement")

	if 4%2 == 0 {
		fmt.Println("4 is even")
	} else {
		fmt.Println("4 is odd")
	}
	name := "pradeep"
	if strings.Contains(name, "pradeep") {
		fmt.Println(name)
	}

	if num := 9; num+1 == 10 {
		fmt.Println(num)
	} else if num == 10 {
		fmt.Println("4 smaller")
	}

	//loops
	fmt.Println("for loop----------------")
	names := []string{"dfsd", "reer", "vvvv", "rree"}

	for i, j := range names {
		fmt.Println(i, j)

	}
	for _, j := range "AAA" {
		fmt.Println(j)
	}

	var a string = "pradee"

	switch a {
	case "pradee":
		fmt.Println("my name", a)
	default:
		fmt.Println("its not my name")
	}

}
