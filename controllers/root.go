package controllers

import (
	// Stdlib
	"net/http"

	// 3rd party
	"github.com/flosch/pongo2"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"

	// Project
	"github.com/danielbarbarito/go-web-template/models"
)

var rootTemplate = pongo2.Must(pongo2.FromFile("templates/index.html"))

type (
	// RootController represents the controller for operating on the Root resource
	RootController struct {
		session *mgo.Session
	}
)

func NewRootController(s *mgo.Session) *RootController {
	return &RootController{s}
}

func (c RootController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u := models.User{}

	u.Name = "Dan Barbarito"

	u.Id = bson.NewObjectId()

	// c.session.DB("test").C("users").Insert(u)

	err := rootTemplate.ExecuteWriter(pongo2.Context{}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
