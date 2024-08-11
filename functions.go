package main

import (
	"fmt"
)

func main() {

	fmt.Println("inside the main funtion")

	//sayhaitoall("Pradee")
	greetingstudents([]string{"aaa", "ddd", "vvv"}, sayhaitoall)
}

func sayhaitoall(s string) {
	fmt.Println("hii there %v !!!!", s)
}
func greetingstudents(s []string, f func(string)) {
	for _, j := range s {
		f(j)
	}

}
