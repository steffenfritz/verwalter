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
	// TODO: Generic path
	tmpl, err := template.ParseFiles(Staticpath + "/templates/addasset.tmpl")
	e(err)
	tmpl.Execute(w, Today.Format(time.RFC3339))
}

// SaveAsset saves new asset entry to database if valid
func SaveAsset(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	aname := r.Form.Get("aname")
	aaddress := r.Form.Get("aaddress")
	ahostname := r.Form.Get("ahostname")
	apurpose := r.Form.Get("apurpose")
	aos := r.Form.Get("aos")
	aosversion := r.Form.Get("aosversion")
	aosupdate := r.Form.Get("aosupdate")
	azone := r.Form.Get("azone")
	aactive := r.Form.Get("aactive")
	avalfrom := r.Form.Get("avalidFrom")
	avalto := r.Form.Get("avalidTo")
	alocation := r.Form.Get("alocation")

	sqlStmt, err := db.Prepare("insert into assets values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, aname, aaddress, ahostname, apurpose, aos, aosversion, aosupdate, azone, aactive, avalfrom, avalto, alocation, 0)
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

	for n := range keys {
		println(keys[n])
	}

	sqlStmt, err := db.Prepare("select * from assets where os='OpenBSD'")
	e(err)

	_, err = sqlStmt.Exec()
	e(err)
}
