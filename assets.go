package main

import (
	"html/template"
	"net/http"
	"time"
)

// Asset is the type that defines an asset
type Asset struct {
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
	responsible  string
	location     string
	functionsID  int
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

// SearchAsset handles requests to addasset
func SearchAsset(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles(Staticpath + "/templates/searchasset.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AssetResult queries the database and prints the result as a list of links that gets by db id
func AssetResult(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	qKeys := map[string]string{"descname": "*", "hostname": "*", "zone": "*"}

	for key, value := range keys {
		if len(value[0]) != 0 {
			qKeys[key] = value[0]
		}
	}

	rows, err := db.Query("select * from assets where descname=? AND hostname=? AND zone=?", qKeys["descname"], qKeys["hostname"], qKeys["zone"])
	e(err)
	defer rows.Close()

	for rows.Next() {
		// NEXT
	}

}
