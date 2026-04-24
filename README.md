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

### Strings

- string literal: one line: ""
- raw string literal: multi lines: ``

```go
package main

import "fmt"

func main() {
	var s string

	s = "now are you?"
	fmt.Println(s)

	s = `now are you?`
	fmt.Println(s)

	s = "<html>\t\n<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	s = `
<html>
	<body>
		"Hello"
	</body>
</html>`
	fmt.Println(s)
}

	fmt.Println("c:\\Code\\Go")
	fmt.Println(`c:\Code\Go`)
```

### UTF-8 length

```go
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
```

### String banger

[Repeat](https://pkg.go.dev/strings#Repeat)
[ToUpper](https://pkg.go.dev/strings#ToUpper)

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)
	fmt.Println(s)

	s = strings.Repeat("!", l) + s
	fmt.Println(s)

}
```

### Strings exercises

#### String ex1

1. Change the following program
2. It should use a raw string literal instead

```go
func main() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"

	path2 := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png
`

	fmt.Println(path)
	fmt.Println(path2)

}
```

#### String ex2

1. Change the following program
2. It should use a raw string literal instead

```go
func main() {
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	json2 := `
{
	"Items": [{
		"Item": {
			"name: "Teddy Bear"
				}
		}]
	}
`

	fmt.Println(json)
	fmt.Println(json2)

}
```

#### Strings ex3
