package main

/* GENERAL NOTE:
Introducing a sec incident to verwalter means introducing a ticketing systems. Which means a lot of work. Therefore postponded.
*/

// Secinc is the type that defines a security incident
type Secinc struct {
	reporterFirstName string
	reporterEmail     string
	reporterTelNo     string
	reportedAsset     int
	reportedService   int
	reportedDate      string
	shortInitDesc     string
	longInitDesc      string
	stillOpen         int
	closedDate        string
	forwardedTo       int
}
