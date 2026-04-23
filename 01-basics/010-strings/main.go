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

	fmt.Println("c:\\Code\\Go")
	fmt.Println(`c:\Code\Go`)

}
