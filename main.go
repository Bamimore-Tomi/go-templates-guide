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
	temp = template.Must(template.ParseFiles("template-06.txt"))
}

func main() {
	// Execute cuteAnimalsSpecies into the template and  print to Stdout
	cuteAnimalsSpecies := map[string]string{
		"Dogs": "German Shepherd",
		"Cats": "Ragdoll",
		"Mice": "Deer Mouse",
		"Fish": "Goldfish",
	}
	err := temp.Execute(os.Stdout, cuteAnimalsSpecies)
	if err != nil {
		log.Fatalln(err)
	}
}

// Animals are cute, some cute animals are:

// Cats , Ragdoll

// Dogs , German Shepherd

// Fish , Goldfish

// Mice , Deer Mouse
