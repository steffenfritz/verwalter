package main

import (
	"html/template"
	"net/http"
)

// Person ist the type that defines a person
type Person struct {
	firstname  string
	middlename string
	lastname   string
	department string
	landline   string
	mobile     string
	street     string
	number     string
	city       string
	zip        string
	country    string
	functions  string
	validFrom  string
	validTo    string
}

// Persons handles requests to persons
func Persons(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/persons.tmpl")
	e(err)
	tmpl.Execute(w, "")
}
