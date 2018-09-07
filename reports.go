package main

import (
	"html/template"
	"net/http"
)

// DefaultReport type is a struct that holds values for the default report
type DefaultReport struct {
	SecIncCountOpen    int
	VulnCountOpen      int
	AssetCountActive   int
	PersonsCountActive int
	OldestUpdate       int
}

// Reports handles requests to reports
func Reports(w http.ResponseWriter, r *http.Request) {
	var DefaultReport DefaultReport

	tmpl, err := template.ParseFiles(Staticpath + "/templates/reports.tmpl")
	e(err)

	rows, err := db.Query("SELECT count(*) FROM secincident where stillOpen = 'true'")
	e(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&DefaultReport.SecIncCountOpen)
	}

	rows, err = db.Query("SELECT count(*) FROM assets where active = 'true'")
	e(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&DefaultReport.AssetCountActive)
	}

	rows, err = db.Query("SELECT count(*) FROM persons")
	e(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&DefaultReport.PersonsCountActive)
	}

	tmpl.Execute(w, DefaultReport)
}
