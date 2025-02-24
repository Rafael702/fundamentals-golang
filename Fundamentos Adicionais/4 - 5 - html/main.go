package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct{
	Nome string
	email string
}

func main() {

	//Pega todo mundo que possui a extensão .html e joga no template
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario {"Joao", "joao.pedro@gmail.com"}


		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na portal 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
