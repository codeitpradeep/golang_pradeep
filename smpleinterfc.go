package main

import "fmt"

type College interface {
	studentdetails() string
}

type stdname struct {
	name string
	dept string
}

type stdbloodgrp struct {
	group string
}

func (n stdname) studentdetails() string {
	return n.name
}

func main() {

	names := stdname{name: "Pradeep", dept: "CSE"}
	//blood := stdbloodgrp{group: "A+"}
	fmt.Println("Interface pgm", callstddetails(names))

}

func callstddetails(s College) string {
	return s.studentdetails()
}
