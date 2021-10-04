package main

import (
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 300
	myNumber := rand.Intn((max-min)+1) + min
	err := temp.Execute(os.Stdout, myNumber)
	if err != nil {
		log.Fatalln(err)
	}
}

// Number 141 is less than 200
