package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
	dept string
	Address
}

type Address struct {
	city  string
	state string
}

func main() {
	//ininltize those nested struct
	address1 := Address{
		city:  "Tiruppur",
		state: "TN",
	}

	//declare inintize value for struct
	employee1 := Employee{
		name:    "Pradeep",
		age:     27,
		dept:    "SRE",
		Address: address1,
	}

	fmt.Println(employee1.name)

	//initilze struct inside main
	job := struct {
		title  string
		salary int
	}{
		title:  "SRE",
		salary: 30000,
	}

	fmt.Println(job.title)

	employee1.updatename("dhanush")

	fmt.Println(employee1)

}

func (e *Employee) updatename(newname string) {
	e.name = newname
}
