package controllers

import (
	"net/http"

	"github.com/sxc/aerialcamp/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

//	type UsersTemplates struct {
//		New views.Template
//	}
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// we need a view to render
	u.Templates.New.Execute(w, nil)
}
