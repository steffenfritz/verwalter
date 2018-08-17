package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var verwalterVersion = "0.1.0"
var dbVersion = "0.1.0"

func connpool(homedir string) {
	dbpath := homedir + "/.verwalter/verwalter.db"

	var err error
	db, err = sql.Open("sqlite3", dbpath)
	e(err)
}

// createDB creates a local database with all tables needed by verwalter
func createDB(homedir string) {

	dbpath := homedir + "/.verwalter/verwalter.db"
	// check if file exists. If error nil, we quit.
	_, err := os.Stat(dbpath)
	if err == nil {
		log.Fatal("Database already exists. Quitting.")
		return
	}

	log.Println("Creating database")
	os.Mkdir(homedir+"/.verwalter", 0700)
	log.Println("Path: " + dbpath)

	db, err := sql.Open("sqlite3", dbpath+"?foreign_keys=ON")
	e(err)
	defer db.Close()

	sqlStmt := `create table basesettings(id INTEGER NOT NULL PRIMARY KEY,
		timezone TEXT,
		lang TEXT,
		processSched TEXT,
		scriptLocation TEXT,
		verwalterVersion TEXT,
		dbVersion TEXT,
		mailServer TEXT,
		mailPort TEXT,
		mailUser TEXT,
		mailPass TEXT
	)`
	_, err = db.Exec(sqlStmt)
	e(err)

	// id is the row id
	// descname is a descriptive name
	// address is a address, e.g. IP or MAC
	// hostname is the assets dns hostname
	// purpose gives a short description why this asset exists, e.g. server, workstation, repo
	// os names the operating system
	// osversion names the version of the operating system
	// lastosupdate gives the date of the last operating system update. Format: yyyy-mm-dd
	// zone names the one where the asset resides, e.g. DMZ
	// active marks an asset as active or not. Due to sqlite3 lack of a BOOL we use INTEGER
	// responsible references a function that is responsible for the host as a service

	sqlStmt = `create table assets(id INTEGER NOT NULL PRIMARY KEY, 
		descname TEXT,
		address TEXT,
		hostname TEXT,
		purpose TEXT,
		os TEXT,
		osversion TEXT,
		lastosupdate TEXT,
		zone TEXT,
		active INTEGER,
		validFrom TEXT,
		validTo TEXT,
		location TEXT,
		responsible INTEGER,
		serialnumber TEXT,
		tagid TEXT,
		assettype INTEGER,
		  FOREIGN KEY (responsible) REFERENCES functions(id),
		  FOREIGN KEY (assettype) REFERENCES assettypes(id)
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// media assets
	sqlStmt = `create table mediaassets(id INTEGER NOT NULL PRIMARY KEY,
		descname TEXT,
		title TEXT,
		creator TEXT,
		author INT,
		lang TEXT,
		location INT,
		contentdesc TEXT,
		  FOREIGN KEY (author) REFERENCES persons(id),
		  FOREIGN KEY (location) REFERENCES location(id)
		);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// location table
	sqlStmt = `create table location(id INTEGER NOT NULL PRIMARY KEY,
		descname TEXT,
		continent TEXT,
		country TEXT,
		state TEXT,
		city TEXT,
		street TEXT,
		number TEXT,
		addition TEXT,
		floor TEXT,
		bureau TEXT,
		room TEXT,
		rack TEXT,
		latitude FLOAT,
		longitude FLOAT		
		);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// assettype table
	sqlStmt = `create table assettypes(id INTEGER NOT NULL PRIMARY KEY,
		assettype TEXT UNIQUE NOT NULL,
		descname TEXT,
		active INTEGER,
		validFrom TEXT,
		validTo TEXT
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// supplier table
	sqlStmt = `create table supplier(id INTEGER NOT NULL PRIMARY KEY,
		suppliername TEXT,
		validFrom TEXT,
		validTo TEXT,
		contact INTEGER,
		  FOREIGN KEY(contact) REFERENCES persons(id)
		);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// services table
	sqlStmt = `create table services(id INTEGER NOT NULL PRIMARY KEY, 
		servicename TEXT,
		application TEXT,
		port INTEGER,
		criticality TEXT,
		license INTEGER,
		  FOREIGN KEY(license) REFERENCES licenses(id)
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// supplierProvides table
	sqlStmt = `create table supplierProvides(id INTEGER NOT NULL PRIMARY KEY,
		supplierID INTEGER,
		serviceID INTEGER,
		  FOREIGN KEY(supplierID) REFERENCES supplier(id),
		  FOREIGN KEY(serviceID) REFERENCES services(id)
		);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// accessedby table contains who can access a service
	sqlStmt = `create table accessedby(id INTEGER NOT NULL PRIMARY KEY,
		serviceid INT NOT NULL,
		assetid INT NOT NULL,
		personid INT NOT NULL,
		  FOREIGN KEY(serviceid) REFERENCES services(id),
		  FOREIGN KEY(assetid) REFERENCES assets(id),
		  FOREIGN KEY(personid) REFERENCES persons(id)
	)`

	// zones is for zone names
	sqlStmt = `create table zones(id INTEGER NOT NULL PRIMARY KEY, 
		name TEXT,
		description TEXT,
		netrange TEXT,
		validFrom TEXT,
		validTo TEXT
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// host_service is a relation table for host to service
	sqlStmt = `create table host_service(id INTEGER NOT NULL PRIMARY KEY, 
		hostid INTEGER,
		serviceid INTEGER,
		active INTEGER,
		validFrom TEXT,
		validTo TEXT,
		  FOREIGN KEY(hostid) REFERENCES assets(id),
		  FOREIGN KEY(serviceid) REFERENCES services(id)
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// reachable_from is a relation table for zones and hosts to host_service
	sqlStmt = `create table reachable_from(id INTEGER NOT NULL PRIMARY KEY,
		zoneid INTEGER,
		hostid INTEGER,
		host_service_id INTEGER,
		validFrom TEXT,
		validTo TEXT,
		  FOREIGN KEY (zoneid) REFERENCES zones(id),
		  FOREIGN KEY (hostid) REFERENCES assets(id),
		  FOREIGN KEY (host_service_id) REFERENCES host_service(id)
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// reaches is a relation table for host_service to zones and hosts
	sqlStmt = `create table reaches(id INTEGER NOT NULL PRIMARY KEY, 
		host_service_id INTEGER,
		zoneid INTEGER,
		hostid INTEGER,
		validFrom TEXT,
		validTo TEXT,
		  FOREIGN KEY (host_service_id) REFERENCES host_service(id),
		  FOREIGN KEY (zoneid) REFERENCES zones(id),
		  FOREIGN KEY (hostid) REFERENCES assets(id)
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// redundancy is a relation table for host_service to host_service
	sqlStmt = `create table redundancy(id INTEGER NOT NULL PRIMARY KEY, 
		host_service_id INTEGER,
		redundant_host_service INTEGER,
		validFrom TEXT,
		validTo TEXT,
		  FOREIGN KEY (host_service_id) REFERENCES host_service(id),
		  FOREIGN KEY (redundant_host_service) REFERENCES host_service(id)
		);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// persons table
	sqlStmt = `create table persons(id INTEGER NOT NULL PRIMARY KEY,
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
		validFrom TEXT,
		validTo TEXT
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// policies table
	sqlStmt = `create table policies(id INTEGER NOT NULL PRIMARY KEY,
		passwordagefactor INT,
		vulnagefactorinfo REAL,
		vulnagefactorlow REAL,
		vulnagefactormedium REAL,
		vulnagefactorhigh REAL,
		secincidentagefactor REAL,
		updateagefactor REAL,
		vulninfofactor REAL,
		vulnlowfactor REAL,
		vulnmediumfactor REAL,
		vulnhighfactor REAL,
		secincidentfactor REAL
		);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// hasaccessto table
	sqlStmt = `create table hasaccessto(id INTEGER NOT NULL PRIMARY KEY,
		personid INT,
		assetid INT,
		serviceid INT,
		validFrom TEXT,
		validTo TEXT,
		  FOREIGN KEY (personid) REFERENCES persons(id),
		  FOREIGN KEY (assetid) REFERENCES assets(id),
		  FOREIGN KEY (serviceid) REFERENCES services(id)
	)`

	// functions table
	sqlStmt = `create table functions(id INTEGER NOT NULL PRIMARY KEY,
		descname TEXT,
		landline TEXT,
		mobile TEXT,
		email TEXT,
		responsibleFirstName TEXT,
		responsibleMiddleName TEXT,
		responsibleLastName TEXT,
		validFrom TEXT,
		validTo TEXT
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// vulnScan table
	sqlStmt = `create table vulnScan(id INTEGER NOT NULL PRIMARY KEY,
		descname TEXT,
		tool TEXT,
		checkdateStart TEXT,
		checkdateStop TEXT
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	// vulns keeps record of all found vulnerable services
	// cve gives the cve identifier
	sqlStmt = `create table vulns(id INTEGER NOT NULL PRIMARY KEY,
		host_service_id INTEGER,
		cve TEXT,
		foundDate TEXT,
		workaroundDate TEXT,
		fixedDate TEXT,
		  FOREIGN KEY (host_service_id) REFERENCES host_service(id)
	);`
	_, err = db.Exec(sqlStmt)
	e(err)

	sqlStmt = `create table downtime(id INTEGER NOT NULL PRIMARY KEY,
		asset_id INTEGER,
		startTime TEXT,
		stopTime TEXT,
		reason TEXT,
		downtime TEXT,
		downatm INTEGER,
		  FOREIGN KEY (asset_id) REFERENCES assets(id)
	)`
	_, err = db.Exec(sqlStmt)
	e(err)

	sqlStmt = `create table os(id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		hyperlink TEXT,
		version TEXT,
		license_id INTEGER,
		  FOREIGN KEY (license_id) REFERENCES licenses(id)
		)`
	_, err = db.Exec(sqlStmt)
	e(err)

	sqlStmt = `create table licenses(id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		hyperlink TEXT,
		version TEXT,
		validFrom TEXT,
		validTo TEXT
	)`
	_, err = db.Exec(sqlStmt)
	e(err)

	sqlStmt = `create table secincident(id INTEGER NOT NULL PRIMARY KEY,
		reporterFirstName TEXT,
		reporterLastName TEXT,
		reporterEmail TEXT,
		reporterTelNo TEXT,
		reportedAsset INTEGER NOT NULL,
		reportedService INTEGER NOT NULL,
		reportedDate TEXT NOT NULL,
		shortInitDesc TEXT,
		longInitDesc TEXT,
		extTicketID TEXT,
		stillOpen INTEGER,
		closedDate TEXT,
			FOREIGN KEY (reportedAsset) REFERENCES assets(id),
			FOREIGN KEY (reportedService) REFERENCES services(id)
	)`
	_, err = db.Exec(sqlStmt)
	e(err)

	// SET DEFAULT VALUES
	// assettype table
	//	sqlStmt = `create table assettypes(id INTEGER NOT NULL PRIMARY KEY,
	//		assettype TEXT UNIQUE NOT NULL,
	//		descname TEXT,
	//		active INTEGER,
	//		validFrom TEXT,
	//		validTo TEXT
	//	)'

	sqlStmt = `insert into assettypes
	(),
	()`

	// DEFAULT ZONE VALUES
	sqlStmt = `insert into zones values
		(null, 'INTERNET', 'Public Internet', '', '', ''),
		(null, 'DMZ', 'Demilitarized Zone', '', '', ''),
		(null, 'INTRANET', 'Intranet', '', '', ''),
		(null, 'PRODUCTION', 'Zone for productive systems', '', '', ''),
		(null, 'INTEGRATION', 'Zone for integrative systems', '', '', ''),
		(null, 'TEST', 'Zone for test systems', '', '', ''),
		(null, 'DEVELOPMENT', 'Zone for development systems', '', '', ''),
		(null, 'AUTHENTICATION', 'Zone for systems related to authentication', '', '', ''),
		(null, 'ADMINISTRATION', 'Zone for administrative and monitoring systems', '', '', '')`

	_, err = db.Exec(sqlStmt)
	e(err)

	// DEFAULT LICENSES
	sqlStmt = `insert into licenses values 
		(null, 'openssh', 'https://cvsweb.openbsd.org/cgi-bin/cvsweb/src/usr.bin/ssh/LICENCE?rev=HEAD', '1.20' ,'2017-04-30', null),
		(null, 'Apache License', 'https://www.apache.org/licenses/LICENSE-2.0.txt', '2.0', '2004-01', null),
		(null, 'X11 License', null, null, null , null),
		(null, 'GPL', 'https://www.gnu.org/licenses/gpl-2.0.en.html', '2.0', '1991-06-02', null),
		(null, 'GPL', 'https://www.gnu.org/licenses/gpl-3.0.en.html', '3.0', '2007-06-29', null),
		(null, 'EPL', 'https://www.eclipse.org/legal/epl-v20.html', '2.0', '2017-08-24', null)`

	_, err = db.Exec(sqlStmt)
	e(err)

	// DEFAULT SERVICES
	sqlStmt = `insert into services values
		(null, 'ssh', 'openssh', 22, 'H', 0),
		(null, 'smtp', 'postfix', 25, 'M', 5),
		(null, 'httpd', 'Apache httpd', 80, 'M', 1),
		(null, 'imap', 'postfix', 143, 'L', 5),
		(null, 'application server', 'Tomcat', 8080, 'M', 1)`

	_, err = db.Exec(sqlStmt)
	e(err)

	// DEFAULT POLICIES
	sql, err := db.Prepare("insert into policies values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	sql.Exec(nil, 90, 30, 14, 7, 3, 1, 14, 0.2, 0.4, 0.6, 1.0, 2.0)
	e(err)

	// BASESETTINGS
	// get local timezone
	tnow := time.Now()
	tzone, _ := tnow.Zone()
	// verwalter language iso 639-1
	lang := "en"
	// processSched defines how often the process scheduler runs in minutes. This must be ignored at update
	processSched := 30
	// default script location
	scriptLocation := ""

	sql, err = db.Prepare("insert into basesettings values(?,?,?,?,?,?,?,?,?,?,?)")
	e(err)

	_, err = sql.Exec(nil, tzone, lang, processSched, scriptLocation, verwalterVersion, dbVersion, "", "", "", "")
	e(err)
}
