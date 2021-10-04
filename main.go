package main

import (
	"log"
	"os"
	"text/template"
)

type PersonTrait struct {
	Name  string
	Trait string
}

// Declare type pointer to a template
var temp *template.Template

// Using the init function to make sure the template is only parsed once in the program
func init() {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	temp = template.Must(template.ParseFiles("template-05.txt"))
}

func main() {
	// Execute person into the template and  print to Stdout
	person := PersonTrait{Name: "Oluwatomisin", Trait: "a great writer"}
	err := temp.Execute(os.Stdout, person)
	if err != nil {
		log.Fatalln(err)
	}

}

// Hello Oluwatomisin

// You are a great writer
