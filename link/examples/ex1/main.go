package main

import (
	"fmt"
	"link"
	"strings"
)

var exampleHtml = `
<html lang="en">
<body>
	<h1>Hello</h1>
	<a href="/other-page">A link to another page</a>
	<a href="/second-page">A link to another another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
