// err.go provides some not so best practice yet handy error sugars
package main

import "log"

// e is short for fatal errors
func e(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

// w is short for non fatal errors, i.e. warnings
func w(err error) {
	if err != nil {
		log.Println(err)
	}
}
