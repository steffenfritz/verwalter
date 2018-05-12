package main

import (
	"flag"
	"os/user"
)

func main() {

	usr, err := user.Current()
	e(err)

	initdb := flag.Bool("init", false, "initialize database")
	flag.Parse()

	if *initdb {
		createDB(usr.HomeDir)
	}
}
