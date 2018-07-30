package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"
)

// Person is the type that defines a person
type Person struct {
	firstname  string
	middlename string
	lastname   string
	department string
	landline   string
	mobile     string
	street     string
	number     string
	city       string
	zip        string
	country    string
	validFrom  string
	validTo    string
}

// SQLPerson is used for sql queries that may return null values
type SQLPerson struct {
	Personid   sql.NullString
	Firstname  sql.NullString
	Middlename sql.NullString
	Lastname   sql.NullString
	Department sql.NullString
	Landline   sql.NullString
	Mobile     sql.NullString
	Street     sql.NullString
	Number     sql.NullString
	City       sql.NullString
	Zip        sql.NullString
	Country    sql.NullString
	ValidFrom  sql.NullString
	ValidTo    sql.NullString
}

// Persons handles requests to persons
func Persons(w http.ResponseWriter, r *http.Request) {
	// TODO: Generic path
	tmpl, err := template.ParseFiles(Staticpath + "/templates/persons.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// AddPerson handles requests to addperson
func AddPerson(w http.ResponseWriter, r *http.Request) {
	Today := time.Now()
	tmpl, err := template.ParseFiles(Staticpath + "/templates/addperson.tmpl")
	e(err)
	tmpl.Execute(w, Today.Format(time.RFC3339))
}

// SavePerson saves new person entry to database
func SavePerson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	e(err)

	p := new(Person)

	p.firstname = r.Form.Get("pfirstname")
	p.middlename = r.Form.Get("pmiddlename")
	p.lastname = r.Form.Get("plastname")
	p.department = r.Form.Get("pdepartment")
	p.landline = r.Form.Get("plandline")
	p.mobile = r.Form.Get("pmobile")
	p.street = r.Form.Get("pstreet")
	p.number = r.Form.Get("pnumber")
	p.city = r.Form.Get("pcity")
	p.zip = r.Form.Get("pzip")
	p.country = r.Form.Get("pcountry")
	p.validFrom = r.Form.Get("pvalidfrom")
	p.validTo = r.Form.Get("pvalidto")

	sqlStmt, err := db.Prepare("insert into persons values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sqlStmt.Exec(nil, p.firstname, p.middlename, p.lastname, p.department, p.landline, p.mobile, p.street, p.number, p.city, p.zip, p.country, p.validFrom, p.validTo)

	Result := ""
	if err != nil {
		Result = "There was an error."
	} else {
		Result = "Person added"
	}
	tmpl, err := template.ParseFiles(Staticpath + "/templates/persons.tmpl")
	e(err)
	tmpl.Execute(w, Result)
}

// SearchPerson handles requests to searchzone
func SearchPerson(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/searchperson.tmpl")
	e(err)
	tmpl.Execute(w, "")
}

// PersonResult queries the database and prints the result as a list of zones that links to all hosts
func PersonResult(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	qKeys := map[string]string{"pfirstname": "%", "plastname": "%", "department": "%"}

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
