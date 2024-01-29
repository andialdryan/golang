package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Andi"
	names[1] = "Muh"
	names[2] = "Aldryan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[2] = "Mandala"
	fmt.Println(names[2])
}