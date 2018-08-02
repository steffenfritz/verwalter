package main

import (
	"html/template"
	"net/http"
)

// TimeFactorPolicy is the record type that holds time factors for calculating escalations and triggers
type TimeFactorPolicy struct {
	PasswordAgeFactor    float64
	VulnAgeFactorInfo    float64
	VulnAgeFactorLow     float64
	VulnAgeFactorMedium  float64
	VulnAgeFactorHigh    float64
	SecincidentAgeFactor float64
	UpdateAgeFactor      float64
}

// SeverityFactor is the record type that holds factors for weighing tasks
type SeverityFactor struct {
	VulnInfo    float64
	VulnLow     float64
	VulnMedium  float64
	VulnHigh    float64
	Secincident float64
}

// Policy is a record type of records TimeFactorPolicy and SeverityFactor
type Policy struct {
	TFP TimeFactorPolicy
	SFP SeverityFactor
}

// Policies handles requests to policies
func Policies(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(Staticpath + "/templates/policies.tmpl")
	e(err)

	var rowid int
	var PolicyStatus Policy
	var TFP TimeFactorPolicy
	var SFP SeverityFactor

	row, err := db.Query("SELECT * FROM policies")
	e(err)
	defer row.Close()

	for row.Next() {
		err = row.Scan(&rowid, &TFP.PasswordAgeFactor, &TFP.VulnAgeFactorInfo, &TFP.VulnAgeFactorLow,
			&TFP.VulnAgeFactorMedium, &TFP.VulnAgeFactorHigh, &TFP.SecincidentAgeFactor, &TFP.UpdateAgeFactor,
			&SFP.VulnInfo, &SFP.VulnLow, &SFP.VulnMedium, &SFP.VulnHigh, &SFP.Secincident)

		e(err)
	}

	PolicyStatus.TFP = TFP
	PolicyStatus.SFP = SFP

	err = tmpl.Execute(w, PolicyStatus)
	e(err)
}

// PasswordEscalation checks the age of passwords and sends notifications if necessary
func PasswordEscalation() {}

// VulnEscalation checks the age and severity of vulnerabilities and sends notifications if necessary
func VulnEscalation() {}

// SecIncEscalation checks the age and severity of security incidents and sends notifications if necessary
func SecIncEscalation() {}
