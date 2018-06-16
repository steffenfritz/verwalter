package main

import (
	"html/template"
	"net/http"
	"time"
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
	validFrom  string
	validTo    string
}

// Persons handles requests to persons
func Persons(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles(Staticpath + "/templates/persons.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AddPerson handles requests to addperson
func AddPerson(w http.ResponseWriter, r *http.Request) {
	Today := time.Now()
	tmpl, err := template.ParseFiles(Staticpath + "/templates/addperson.tmpl")
	e(err)
	tmpl.Execute(w, Today.Format(time.RFC3339))
}

// SavePerson saves new person entry to database
func SavePerson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	p := new(Person)

	p.firstname = r.Form.Get("pfirstname")
	p.middlename = r.Form.Get("pmiddlename")
	p.lastname = r.Form.Get("plastname")
	p.department = r.Form.Get("pdepartment")
	p.landline = r.Form.Get("plandline")
	p.mobile = r.Form.Get("pmobile")
	p.street = r.Form.Get("pstreet")
	p.number = r.Form.Get("pnumber")
	p.city = r.Form.Get("pcity")
	p.zip = r.Form.Get("pzip")
	p.country = r.Form.Get("pcountry")
	p.validFrom = r.Form.Get("pvalidfrom")
	p.validTo = r.Form.Get("pvalidto")

	sqlStmt, err := db.Prepare("insert into persons values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, p.firstname, p.middlename, p.lastname, p.department, p.landline, p.mobile, p.street, p.number, p.city, p.zip, p.country, p.validFrom, p.validTo)

	Result := ""
	if err != nil {
		Result = "There was an error."
	} else {
		Result = "Person added"
	}
	tmpl, err := template.ParseFiles(Staticpath + "/templates/persons.tmpl")
	e(err)
	tmpl.Execute(w, Result)
}
