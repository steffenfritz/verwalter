package main

import (
	"html/template"
	"net/http"
	"time"
)

// Secinc is the type that defines a security incident
// verwalter is not a ticketing tool, therefore Secinc only
// holds some relevant info and a link to an external ticketing tool
type Secinc struct {
	reporterFirstName string
	reporterLastName  string
	reporterEmail     string
	reporterTelNo     string
	reportedAsset     string
	reportedService   string
	reportedDate      string
	shortInitDesc     string
	longInitDesc      string
	extTicketID       string
	stillOpen         string
	closedDate        string
}

// Secincident handles requests to secincident
func Secincident(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/secincident.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AddSecincident handles requests to addperson
func AddSecincident(w http.ResponseWriter, r *http.Request) {
	Today := time.Now()
	tmpl, err := template.ParseFiles(Staticpath + "/templates/addsecincident.tmpl")
	e(err)
	tmpl.Execute(w, Today.Format(time.RFC3339))
}

// SaveSecincident handles requests to addperson
func SaveSecincident(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	s := new(Secinc)

	s.reporterFirstName = r.Form.Get("sifirstname")
	s.reporterLastName = r.Form.Get("silastname")
	s.reporterEmail = r.Form.Get("siemail")
	s.reporterTelNo = r.Form.Get("sitelno")
	s.reportedAsset = r.Form.Get("siasset")
	s.reportedService = r.Form.Get("siservice")
	s.reportedDate = r.Form.Get("sidate")
	s.shortInitDesc = r.Form.Get("sishortinitdesc")
	s.longInitDesc = r.Form.Get("silonginitdesc")
	s.extTicketID = r.Form.Get("siextticketid")
	s.stillOpen = r.Form.Get("sistillopen")
	s.closedDate = r.Form.Get("sicloseddate")

	sqlStmt, err := db.Prepare("insert into secincident values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, s.reporterFirstName, s.reporterLastName, s.reporterEmail, s.reporterTelNo, s.reportedAsset, s.reportedService, s.reportedDate, s.shortInitDesc, s.longInitDesc, s.extTicketID, s.stillOpen, s.closedDate)

	Result := ""
	if err != nil {
		Result = "There was an error."
	} else {
		Result = "Incident added"
	}
	tmpl, err := template.ParseFiles(Staticpath + "/templates/secincident.tmpl")
	e(err)
	tmpl.Execute(w, Result)

}

// SearchSecincident searches security incidents
func SearchSecincident(w http.ResponseWriter, r *http.Request) {}

// SecincidentResults queries the database for security incidents
func SecincidentResult(w http.ResponseWriter, r *http.Request) {}
