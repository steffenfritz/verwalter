package main

import (
	"database/sql"
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

// SQLSecinc is used for sql queries that may return null values
type SQLSecinc struct {
	Secid             sql.NullString
	ReporterFirstName sql.NullString
	ReporterLastName  sql.NullString
	ReporterEmail     sql.NullString
	ReporterTelNo     sql.NullString
	ReportedAsset     sql.NullString
	ReportedService   sql.NullString
	ReportedDate      sql.NullString
	ShortInitDesc     sql.NullString
	LongInitDesc      sql.NullString
	ExtTicketID       sql.NullString
	StillOpen         sql.NullString
	ClosedDate        sql.NullString
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
		Result = "Person added"
	}
	tmpl, err := template.ParseFiles(Staticpath + "/templates/secincident.tmpl")
	e(err)
	tmpl.Execute(w, Result)

}

// SearchSecincident handles requests to searchzone
func SearchSecincident(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/searchsecincident.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// SecincidentResult queries the database and prints the result as a list of sec incidents
func SecincidentResult(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	qKeys := map[string]string{"": "%", "plastname": "%", "department": "%"}

	for key, value := range keys {
		if len(value[0]) != 0 {
			qKeys[key] = value[0]
		}
	}

	rows, err := db.Query("SELECT * FROM persons WHERE (COALESCE(firstname, '') LIKE ?) AND (COALESCE(lastname, '') LIKE ?) AND (COALESCE(department,'') LIKE ?)", qKeys["pfirstname"], qKeys["plastname"], qKeys["department"])
	e(err)
	defer rows.Close()

	var ResultList []SQLPerson
	for rows.Next() {
		var tempResult SQLPerson
		err := rows.Scan(&tempResult.Personid, &tempResult.Firstname, &tempResult.Middlename,
			&tempResult.Lastname, &tempResult.Department, &tempResult.Landline, &tempResult.Mobile,
			&tempResult.Street, &tempResult.Number, &tempResult.City, &tempResult.Zip,
			&tempResult.Country, &tempResult.ValidFrom, &tempResult.ValidTo)

		e(err)

		ResultList = append(ResultList, tempResult)
	}

	tmpl, err := template.ParseFiles(Staticpath + "/templates/resultperson.tmpl")
	e(err)
	err = tmpl.Execute(w, ResultList)
	e(err)
}
