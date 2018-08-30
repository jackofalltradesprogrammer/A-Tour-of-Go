package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s = i.(string) // statement assert if i holds the concrete type string
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Printf(f)
}
