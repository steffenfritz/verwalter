package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func serv() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/assets", Assets)
	r.HandleFunc("/addasset", AddAsset)
	r.HandleFunc("/saveasset", SaveAsset)
	r.HandleFunc("/searchasset", SearchAsset)
	r.HandleFunc("/networks", Networks)
	r.HandleFunc("/persons", Persons)
	r.HandleFunc("/policies", Policies)
	r.HandleFunc("/vulns", Vulns)
	r.HandleFunc("/secincident", Secincident)
	r.HandleFunc("/processes", Processes)
	r.HandleFunc("/reports", Reports)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8666",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

// HomeHandler handles requests to root
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/index.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Assets handles requests to assets
func Assets(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/assets.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AddAsset handles requests to addasset
func AddAsset(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/addasset.tmpl")
	e(err)
	tmpl.Execute(w, "")
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
	aactive := r.Form.Get("azone")

	sqlStmt, err := db.Prepare("insert into assets values(?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, aname, aaddress, ahostname, apurpose, aos, aosversion, aosupdate, azone, aactive, 0)
	e(err)

	result := ""
	if err != nil {
		result = "There was an error."
	} else {
		result = "Asset added"
	}
	tmpl, err := template.ParseFiles("templates/assets.tmpl")
	e(err)
	tmpl.Execute(w, result)
}

// SearchAsset handles requests to addasset
func SearchAsset(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/searchasset.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AssetResult queries the database and prints the result as a list of links that gets by db id
func AssetResult(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	//aname := r.Form.Get("aname")
	//aaddress := r.Form.Get("aaddress")
	//ahostname := r.Form.Get("ahostname")
	//apurpose := r.Form.Get("apurpose")
	aos := r.Form.Get("aos")
	//aosversion := r.Form.Get("aosversion")
	//aosupdate := r.Form.Get("aosupdate")
	//azone := r.Form.Get("azone")
	//aactive := r.Form.Get("azone")

	searchTerm := "aos = " + aos

	sqlStmt, err := db.Prepare("select * from assets where " + searchTerm)
	e(err)

	_, err = sqlStmt.Exec()
	e(err)
}

// ShowAsset shows a single asset entry
func ShowAsset(w http.ResponseWriter, r *http.Request) {}

// Networks handles requests to networks
func Networks(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/networks.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Persons handles requests to persons
func Persons(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/persons.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Policies handles requests to policies
func Policies(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/policies.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Vulns handles requests to vulns
func Vulns(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/vulnerables.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Secincident handles requests to secincident
func Secincident(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/secincident.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Processes handles requests to processes
func Processes(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/processes.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Reports handles requests to reports
func Reports(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/reports.tmpl")
	e(err)
	tmpl.Execute(w, "")
}
