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
	r.HandleFunc("/assetresult", AssetResult)
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

// Networks handles requests to networks
func Networks(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles("templates/networks.tmpl")
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
