package main

import "fmt"

func main() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		firday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, firday, saturday, sunday)

	const (
		EST = -(iota + 5)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}
