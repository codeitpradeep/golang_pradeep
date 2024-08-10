package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	msg := "Heelo im pradeep"

	fmt.Println(strings.Contains(msg, "pradeep"))
	fmt.Println(strings.ReplaceAll(msg, "pradeep", "me"))
	fmt.Println(strings.Split(msg, " "))
	fmt.Println(strings.Index(msg, "p"))

	ages := []int{23, 54, 43, 66, 22, 30}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	name := []string{"dsadsd", "sdfdsf", "wesd", "ygdf", "acvsrf"}

	sort.Strings(name)

	fmt.Println(name)

	fmt.Println(sort.SearchStrings(name, "wesd"))
}

/*
true
Heelo im me
[Heelo im pradeep]
9
[22 23 30 43 54 66]
2
[acvsrf dsadsd sdfdsf wesd ygdf]
3
*/
