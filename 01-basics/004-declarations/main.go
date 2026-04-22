package main

import "fmt"

var i = 10

func shadowing() {
	fmt.Println("package's i:", i)

	// package's i is being shadowed:
	i := 5
	// i above is a new variable.

	fmt.Println("main block:", i)

	// a new scope begins
	{
		// i is a new variable
		i := 6
		fmt.Println("inner block:", i)
	}
	// the scope ends

	// main's i
	fmt.Println("main's i:", i)
}

func main() {
	var area int // zero value: 0

	width := 100
	heigh := 200

	area = width * heigh

	fmt.Println(area)

	shadowing()
}
