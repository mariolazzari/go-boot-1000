# Go Bootcamp: Master Golang with 1000+ Exercises and Projects

[Udemy](https://www.udemy.com/course/learn-go-the-complete-bootcamp-course-golang/learn/lecture/14762174#overview)
[Github](https://github.com/inancgumus/learngo/)

## Basics

### Installation

[YouTube](https://www.youtube.com/watch?v=1MXIGYrMk80)

### Variables

[Hoisting](https://jsbin.com/kilumofuxu/edit?js,console,output)

### Path separators

[split](https://pkg.go.dev/path#Split)
[blank assign](https://go.dev/doc/effective_go#blank_assign)

### Declarations

#### Use var

- if you don't know the initial value
- at package scope
- group multiple vars

#### Short declaration

- most used
- init values known
- redeclaration

### Convert values

```go
package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	speed = speed * int(force)
	fmt.Println(speed)

	speed = 100
	speed = int(float64(speed) * force)
	fmt.Println(speed)

}
```

### Terminal input

[os](https://pkg.go.dev/os)
[Slices](https://go.dev/tour/moretypes/7)

```sh
go build -o greeter
./greeter ciao mario
```

```go
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
```

### Naming conventions

- use first few letters
- use complete words in larger scope
- use mixed caps
- use capital letters for acronym
- do not use underscores
