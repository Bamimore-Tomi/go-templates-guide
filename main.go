package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// Declare type pointer to a template
var temp *template.Template

// Using the init function to make sure the template is only parsed once in the program
func init() {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	temp = template.Must(template.New("template-08.txt").Funcs(funcMap).ParseFiles("template-08.txt"))
}

// Custom function must have only 1 return value, or 1 return value and an error
func formatDate(timeStamp time.Time) string {
	//Define layout for formatting timestamp to string
	return timeStamp.Format("01-02-2006")
}

// Map name formatDate to formatDate function above
var funcMap = template.FuncMap{
	"formatDate": formatDate,
}

func main() {
	timeNow := time.Now()
	err := temp.Execute(os.Stdout, timeNow)
	if err != nil {
		log.Fatalln(err)
	}
}

// Hi,

// Time before formatting : 2021-10-04 18:01:59.6659258 +0100 WAT m=+0.004952101
// Time After formatting : 09-04-2021
