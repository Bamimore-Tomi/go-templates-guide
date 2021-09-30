package main

import (
	"log"
	"os"
	"text/template"
)

// Parse all the files in a certain directory
func main() {
	// This function takes a pattern. It can be a folder
	temp, err := template.ParseGlob("templates/*")

	if err != nil {
		log.Fatalln(err)
	}
	// Simply calling execute parses the first file in the directory
	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// Or you can execute a particular template in the directory
	err = temp.ExecuteTemplate(os.Stdout, "template-03.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// Calling a template not in the directory will produce an error
	err = temp.ExecuteTemplate(os.Stdout, "template-04.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
