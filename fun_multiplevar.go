package main

import (
	"fmt"
	"strings"
)

func main() {
	v1, v2 := getInitials("Pradeep vijay")

	fmt.Println(v1, v2)

}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initvar []string
	for _, v := range names {
		initvar = append(initvar, v[:1])
	}

	if len(initvar) > 1 {
		return initvar[0], initvar[1]
	}

	return initvar[0], "_"
}
