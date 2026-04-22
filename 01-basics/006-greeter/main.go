package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Printf("Path: %s\n", os.Args[0])
	fmt.Printf("Total args: %d\n", len(os.Args))
}
