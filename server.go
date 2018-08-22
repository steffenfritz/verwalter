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
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(Staticpath+"/static"))))
	r.HandleFunc("/assets", Assets)
	r.HandleFunc("/addasset", AddAsset)
	r.HandleFunc("/saveasset", SaveAsset)
	r.HandleFunc("/searchasset", SearchAsset)
	r.HandleFunc("/assetresult", AssetResult)
	r.HandleFunc("/zones", Zones)
	r.HandleFunc("/addzone", AddZone)
	r.HandleFunc("/savezone", SaveZone)
	r.HandleFunc("/searchzone", SearchZone)
	r.HandleFunc("/zoneresult", ZoneResult)
	r.HandleFunc("/persons", Persons)
	r.HandleFunc("/addperson", AddPerson)
	r.HandleFunc("/saveperson", SavePerson)
	r.HandleFunc("/searchperson", SearchPerson)
	r.HandleFunc("/personresult", PersonResult)
	r.HandleFunc("/policies", Policies)
	r.HandleFunc("/vulns", Vulns)
	r.HandleFunc("/secincident", Secincident)
	r.HandleFunc("/addsecincident", AddSecincident)
	r.HandleFunc("/savesecincident", SaveSecincident)
	r.HandleFunc("/searchsecincident", SearchSecincident)
	r.HandleFunc("/secincresult", SecincidentResult)
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
	tmpl, err := template.ParseFiles(Staticpath + "/templates/index.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// All the handlers below have be refactored into their resp. source files

// Vulns handles requests to vulns
func Vulns(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/vulnerables.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// Processes handles requests to processes
func Processes(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/processes.tmpl")
	e(err)
	tmpl.Execute(w, "")
}
