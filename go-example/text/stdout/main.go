package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// <!DOCTYPE html>
// <html lang="en">
// <head>
//     <meta charset="UTF-8">
//     <title>Hello World!</title>
// </head>
// <body>
// <h1>Hello</h1>
// </body>
