package main

import (
	"log"
	"net/http"
	secrets "store-app/src/config"
	"store-app/src/routes"

	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	n.UseHandler(routes.RegisterRoutes())

	err := http.ListenAndServe(":"+secrets.GetSecrets().ApplicationPort, n)
	if err != nil {
		log.Fatal(err)

		if r := recover(); r != nil {
			err = r.(error)
		}
	}
}
