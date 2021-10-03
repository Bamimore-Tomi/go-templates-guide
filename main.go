package main

import (
	"log"
	"os"
	"text/template"
)

// Declare type pointer to a template
var temp *template.Template

// Using the init function to make sure the template is only parsed once in the program
func init() {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	temp = template.Must(template.ParseFiles("template-04.txt"))
}

func main() {
	// Execute myName into the template and  print to Stdout
	myName := "Oluwatomisin"
	err := temp.Execute(os.Stdout, myName)
	if err != nil {
		log.Fatalln(err)
	}

}

// Hello Oluwatomisin

// You are doing great, keep learning.
// Do not stop Oluwatomisin
