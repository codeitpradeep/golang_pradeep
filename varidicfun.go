package main

import (
	"fmt"
)

func callvaridic(s ...string) {

	fmt.Println(s)
	name := ""

	for _, num := range s {
		name += num
	}

	fmt.Println(name)
}

func main() {
	fmt.Println("Hi......")

	v := []string{"a", "b", "c", "d", "e", "f", "i"}

	fmt.Println(v)
	callvaridic(v...)

}
