package main

// TimeFactorPolicy is the record type that holds time factors for calculating escalation trigger
type TimeFactorPolicy struct {
	PasswordAgeFactor    int
	VulnAgeFactorInfo    int
	VulnAgeFactorLow     int
	VulnAgeFactorMedium  int
	VulnAgeFactorHigh    int
	SecincidentAgeFactor int
	UpdateAgeFactor      int
}

// SeverityFactor is the record type that holds factors for weighing tasks
type SeverityFactor struct {
	VulnInfo    float64
	VulnLow     float64
	VulnMedium  float64
	VulnHigh    float64
	Secincident float64
}
