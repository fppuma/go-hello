package main

import "fmt"

func main() {
	var i, j int = 1, 2

	//Inside a function, the := short assignment statement can be used
	//in place of a var declaration with implicit type.
	k := 3

	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
