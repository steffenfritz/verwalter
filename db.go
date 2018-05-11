// db.go is the source file for database related stuff
package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// createDB creates a local database with all tables needed by verwalter
func createDB(filename string) {

	// check if file exists. If error nil, we quit.
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatal("Database already exists. Quitting.")
		return
	}

	log.Println("Creating database")
	log.Println("Path: " + filename)

	db, err := sql.Open("sqlite", filename)
	e(err)
	defer db.Close()

	// id is the row id
	// descname is a descriptive name
	// address is a address, e.g. IP or MAC
	// hostname is the assets dns hostname
	// purpose gives a short description why this asset exists, e.g. server, workstation, repo
	// os names the operating system
	// osversion names the version of the operating system
	// lastosupdate gives the date of the last operating system update. Format: yyyy-mm-dd
	// zone names the one where the asset resides, e.g. DMZ
	// services lists services on the host by name or port and proto, e.g. http or 443/tcp
	// reachableFrom lists zones from where a service is reachable as a tuple, e.g. (internet, http) or (intern, tcp/443)
	// reaches lists which zones/hosts and services the host can reach, e.g. (internet, http) or (10.0.1.1, tcp/443)
	// active marks an asset as active or not. Due to sqlite3 lack of a BOOL we use INTEGER
	// vulnerable lists vulnerable packages/services. If services or packages are vulnerable, they are listed here.
	//   So if not null, host is vulnerable.
	// redundancy lists hosts that are redundant to the host. If null, host and its services have no redundancy
	// responsibles lists functions ids that are responsible for the host and its services

	sqlStmt := `create table assets(id INTEGER NOT NULL PRIMARY KEY, 
		descname TEXT,
		address TEXT,
		hostname TEXT,
		purpose TEXT,
		os TEXT,
		osversion TEXT,
		lastosupdate TEXT,
		zone TEXT,
		services TEXT,
		reachableFrom TEXT,
		reaches TEXT,
		active INTEGER,
		vulnerable TEXT,
		redundancy TEXT,
		responsibles TEXT
	);`
	_, err = db.Exec(sqlStmt)

	// NOTE: If something changed in the persons's dates,
	//       a new entry for the same person is created
	//
	// id is the row id
	// most fields are self-descriptive
	// functions lists functions of the person the ITIL way, e.g. [SAP, HELPDESK]
	// validFrom gives the date from when the person was active with this data
	// validTo gives the date until the person was active with this data
	// wasID gives the ID of the same person before something has changed in its dates
	// becameID gives the ID of the same person after something has changed in its dates
	// hasAccessTo lists asset ID-service tuples the person has access to, e.g. ()
	sqlStmt = `create table persons (id INTEGER NOT NULL PRIMARY KEY,
		firstname TEXT,
		middlename TEXT,
		lastname TEXT,
		department TEXT,
		landline TEXT,
		mobile TEXT,
		street TEXT,
		number TEXT,
		city TEXT,
		zip TEXT,
		country TEXT,
		functions TEXT,
		validFrom TEXT,
		validTo TEXT,
		wasID INTEGER,
		hasAccessTo TEXT
	);`

	_, err = db.Exec(sqlStmt)

	// id is the row id
	// descname is a descriptive name
	// responsibleName gives the name of a responsible person for the function
	// most fields are self-descriptive

	sqlStmt = `create table functions (id INTEGER NOT NULL PRIMARY KEY,
		descname TEXT,
		landline TEXT,
		mobile TEXT,
		email TEXT,
		responsibleFirstName TEXT,
		responsibleMiddleName TEXT,
		responsibleLastName TEXT,
		validFrom TEXT,
		validTo TEXT,
	);`
}
