package main

import (
	"flag"
	"log"
	"os"
	"os/user"
)

// Staticpath sets the path to additional static ressources like css and templates
var Staticpath string

func main() {

	usr, err := user.Current()
	e(err)

	initdb := flag.Bool("init", false, "initialize database")
	flag.Parse()

	Staticpath = usr.HomeDir + "/.verwalter"

	if *initdb {
		createDB(usr.HomeDir)
		log.Println("Created database.")
		log.Println("You can now start the application by starting 'verwalter'")
		os.Exit(0)
	}

	connpool(usr.HomeDir)

	log.Println("Please visit http://127.0.0.1:8666")
	log.Println("Trying to open your default browser")
	openbrowser("http://127.0.0.1:8666")

	serv()
	defer db.Close()
}
