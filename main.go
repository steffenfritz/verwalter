package main

import (
	"flag"
	"log"
	"os"
	"os/user"
)

func main() {

	usr, err := user.Current()
	e(err)

	initdb := flag.Bool("init", false, "initialize database")
	flag.Parse()

	if *initdb {
		createDB(usr.HomeDir)
		log.Println("Created database.")
		log.Println("You can now start the application by starting 'verwalter'")
		os.Exit(0)
	}

	log.Println("Please visit http://127.0.0.1:8666")
	log.Println("Trying to open your default browser")
	openbrowser("http://127.0.0.1:8666")
	serv()

}
