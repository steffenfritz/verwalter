package main

import (
	"database/sql"
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

// SQLZone is used to unmarshal sql queries with posible nul values
type SQLZone struct {
	Zoneid      sql.NullString
	Name        sql.NullString
	Description sql.NullString
	Netrange    sql.NullString
	ValidFrom   sql.NullString
	ValidTo     sql.NullString
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

// SearchZone handles requests to searchzone
func SearchZone(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/searchzone.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// ZoneResult queries the database and prints the result as a list of zones that links to all hosts
func ZoneResult(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	qKeys := map[string]string{"zname": "%", "zdesc": "%"}

	for key, value := range keys {
		if len(value[0]) != 0 {
			qKeys[key] = value[0]
		}
	}

	rows, err := db.Query("SELECT * FROM zones WHERE (COALESCE(name, '') LIKE ?) AND (COALESCE(description, '') LIKE ?)", qKeys["zname"], qKeys["zdesc"])
	e(err)
	defer rows.Close()

	var ResultList []SQLZone
	for rows.Next() {
		var tempResult SQLZone
		err := rows.Scan(&tempResult.Zoneid, &tempResult.Name, &tempResult.Description, &tempResult.Netrange, &tempResult.ValidFrom, &tempResult.ValidTo)

		e(err)

		ResultList = append(ResultList, tempResult)
	}

	tmpl, err := template.ParseFiles(Staticpath + "/templates/resultzones.tmpl")
	e(err)
	err = tmpl.Execute(w, ResultList)
	e(err)
}
