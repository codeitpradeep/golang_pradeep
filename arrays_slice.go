package main

import "fmt"

func main() {

	var args = [3]int{1, 2, 3}

	fmt.Println(args, len(args))

	name := [4]string{"aaa", "vvv", "eee", "ffff"}

	fmt.Println(name, len(name))

	name[2] = "llllll"
	//append(name, "ooooooo")

	fmt.Println(name)

	var school = []string{}

	school = append(school, "p")

	fmt.Println(school)

	var vowls = []string{"a", "e", "i", "o", "u"}
	fmt.Println(vowls)
	r1 := vowls[1:4]
	r2 := vowls[2:]
	fmt.Println("range1 [1:4] ", r1)
	fmt.Println("range2 [2:] ", r2)

	vowls = append(r2, "aaaa")

	fmt.Println(vowls)
}

/*
[1 2 3] 3
[aaa vvv eee ffff] 4
[aaa vvv llllll ffff]
[p]
[a e i o u]
range1 [1:4]  [e i o]
range2 [2:]  [i o u]
[i o u aaaa]
*/
