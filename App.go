package main

import (
	// Stdlib
	"net/http"

	// 3rd party
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	// Project
	"github.com/danielbarbarito/go-web-template/controllers"
)

func getSession() *mgo.Session {
		// Connect to our local mongo
		s, err := mgo.Dial("mongodb://localhost")

		// Check if connection error, is mongo running?
		if err != nil {
				panic(err)
		}

		return s
}


func main() {

	router := httprouter.New()

	// Routes
	rootController := controllers.NewRootController(getSession())
	router.GET("/", rootController.Show)

	// Static file server
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	http.ListenAndServe(":3000", router)
}
