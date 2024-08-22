package main

import "fmt"

func main() {

	fmt.Println("--map--")

	m := make(map[string]int)

	m["AA"] = 1
	m["BB"] = 2
	m["CC"] = 3

	for i, v := range m {
		fmt.Println(i, "-", v)
	}

	names := map[int]string{
		1: "AAA",
		2: "BBB",
		3: "CCC",
		4: "DDD",
		5: "EEE",
	}

	fmt.Println(names)

	delete(names, 2)

	fmt.Println("After Delete-", names)

	clear(names)

	fmt.Println(names)
}
