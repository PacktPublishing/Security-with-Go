package main

import (
	"fmt"
	"html"
)

func main() {
	rawString := `<script>alert("Test");</script>`
	safeString := html.EscapeString(rawString)

	fmt.Println("Unescaped: " + rawString)
	fmt.Println("Escaped: " + safeString)
}