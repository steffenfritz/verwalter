package main

import (
	"html/template"
	"net/http"
	"time"
)

// Zone defines a network zone
type Zone struct {
	name        string
	description string
	netrange    string
	validFrom   string
	validTo     string
}

// Zones handles requests to zones
func Zones(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/zones.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AddZone handles requests to addzone
func AddZone(w http.ResponseWriter, r *http.Request) {
	Today := time.Now()
	tmpl, err := template.ParseFiles(Staticpath + "/templates/addzone.tmpl")
	e(err)
	tmpl.Execute(w, Today.Format(time.RFC3339))
}

// SaveZone saves new zone entry to database if valid
func SaveZone(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	z := new(Zone)

	z.name = r.Form.Get("zname")
	z.description = r.Form.Get("zdescription")
	z.netrange = r.Form.Get("znetrange")
	z.validFrom = r.Form.Get("zvalidFrom")
	z.validTo = r.Form.Get("zvalidTo")

	sqlStmt, err := db.Prepare("insert into zones values(?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, z.name, z.description, z.netrange, z.validFrom, z.validTo)
	e(err)

	Result := ""
	if err != nil {
		Result = "There was an error."
	} else {
		Result = "Zone added"
	}
	tmpl, err := template.ParseFiles(Staticpath + "/templates/zones.tmpl")
	e(err)
	tmpl.Execute(w, Result)
}
