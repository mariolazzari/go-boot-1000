package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "mario"
	fmt.Println(len(name))

	name = "Mariò"
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))
}
