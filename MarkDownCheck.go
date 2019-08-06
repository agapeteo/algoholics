package main

import (
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
)

func not_main () {
	fmt.Println("Yo")
	text := join("*one*", " #two#", " ```three```")
	input := []byte(text)
	output := blackfriday.Run(input)
	fmt.Println(string(output))
}

