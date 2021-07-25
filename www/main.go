package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handleRequest()
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	//fmt.Fprintf(w, "Go is super easy!!!")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, happiness float64
}
