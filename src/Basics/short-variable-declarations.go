package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // not available outside the function
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
