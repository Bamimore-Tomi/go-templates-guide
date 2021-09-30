package main

import (
	"log"
	"os"
	"text/template"
)

// Prints out the template without passing any value using the text/template package
func main() {
	template, err := template.ParseFiles("template-01.txt")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}
	// Print out the template to std
	template.Execute(os.Stdout, nil)
}

//OUTPUT
// Hi <no value>

// You are are welcome to this tutorial
