package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"
)

// Asset is the type that defines an asset
type Asset struct {
	assetid      string
	descname     string
	address      string
	hostname     string
	purpose      string
	os           string
	osversion    string
	lastosupdate string
	zone         string
	active       string
	validFrom    string
	validTo      string
	location     string
	responsible  int
}

// SQLAsset is used to unmarshal sql queries with possible null values
type SQLAsset struct {
	Assetid      sql.NullString
	Descname     sql.NullString
	address      sql.NullString
	hostname     sql.NullString
	purpose      sql.NullString
	os           sql.NullString
	osversion    sql.NullString
	lastosupdate sql.NullString
	zone         sql.NullString
	active       sql.NullString
	validFrom    sql.NullString
	validTo      sql.NullString
	location     sql.NullString
	responsible  sql.NullInt64
}

// Assets handles requests to assets
func Assets(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/assets.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AddAsset handles requests to addasset
func AddAsset(w http.ResponseWriter, r *http.Request) {
	Today := time.Now()
	tmpl, err := template.ParseFiles(Staticpath + "/templates/addasset.tmpl")
	e(err)
	tmpl.Execute(w, Today.Format(time.RFC3339))
}

// SaveAsset saves new asset entry to database if valid
func SaveAsset(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	a := new(Asset)

	a.descname = r.Form.Get("aname")
	a.address = r.Form.Get("aaddress")
	a.hostname = r.Form.Get("ahostname")
	a.purpose = r.Form.Get("apurpose")
	a.os = r.Form.Get("aos")
	a.osversion = r.Form.Get("aosversion")
	a.lastosupdate = r.Form.Get("aosupdate")
	a.zone = r.Form.Get("azone")
	a.active = r.Form.Get("aactive")
	a.validFrom = r.Form.Get("avalidFrom")
	a.validTo = r.Form.Get("avalidTo")
	a.location = r.Form.Get("alocation")

	sqlStmt, err := db.Prepare("insert into assets values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, a.descname, a.address, a.hostname, a.purpose, a.os, a.osversion, a.lastosupdate, a.zone, a.active, a.validFrom, a.validTo, a.location, 0)
	e(err)

	Result := ""
	if err != nil {
		Result = "There was an error."
	} else {
		Result = "Asset added"
	}
	tmpl, err := template.ParseFiles(Staticpath + "/templates/assets.tmpl")
	e(err)
	tmpl.Execute(w, Result)
}

// SearchAsset handles requests to searchasset
func SearchAsset(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/searchasset.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AssetResult queries the database and prints the result as a list of links that gets by db id
func AssetResult(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	qKeys := map[string]string{"descname": "%", "ahostname": "%", "azone": "%"}

	for key, value := range keys {
		if len(value[0]) != 0 {
			qKeys[key] = value[0]
		}
	}

	rows, err := db.Query("SELECT * FROM assets WHERE (COALESCE(descname, '') LIKE ?) AND (COALESCE(hostname, '') LIKE ?) AND (COALESCE(zone,'') LIKE ?)", qKeys["descname"], qKeys["ahostname"], qKeys["azone"])
	e(err)
	defer rows.Close()

	var ResultList []SQLAsset
	for rows.Next() {
		var tempResult SQLAsset
		err := rows.Scan(&tempResult.Assetid, &tempResult.Descname, &tempResult.address, &tempResult.hostname,
			&tempResult.purpose, &tempResult.os, &tempResult.osversion, &tempResult.lastosupdate,
			&tempResult.zone, &tempResult.active, &tempResult.validFrom, &tempResult.validTo,
			&tempResult.location, &tempResult.responsible)

		e(err)

		ResultList = append(ResultList, tempResult)
	}

	tmpl, err := template.ParseFiles(Staticpath + "/templates/resultassets.tmpl")
	e(err)
	err = tmpl.Execute(w, ResultList)
	e(err)
}
