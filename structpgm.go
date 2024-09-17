package main

import "fmt"

type collage struct {
	name    string
	rollnum int
	dept    string
}

func main() {
	addstudent("pradeep", 122, "CSE")
}

func addstudent(s1 string, i int, s2 string) {
	fmt.Println(collage{s1, i, s2})

}
