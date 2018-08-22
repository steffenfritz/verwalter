package main

import (
	"html/template"
	"net/http"
)

// Reports handles requests to reports
func Reports(w http.ResponseWriter, r *http.Request) {
	var SecIncCountOpen int

	tmpl, err := template.ParseFiles(Staticpath + "/templates/reports.tmpl")
	e(err)

	rows, err := db.Query("SELECT count(*) FROM secincident where stillOpen = 'true'")
	e(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&SecIncCountOpen)
	}

	tmpl.Execute(w, SecIncCountOpen)
}
