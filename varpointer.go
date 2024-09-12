package main

import "fmt"

func main() {

	name := "pradeep"
	refanceofit := &name
	fmt.Println(refanceofit, " ", name)

	name = "aaa"
	fmt.Println(name)

	*refanceofit = "arun"
	fmt.Println(name)

	getname(refanceofit)

	fmt.Println(name)

}

func getname(refanceofit *string) {
	*refanceofit = "vijay"
}
