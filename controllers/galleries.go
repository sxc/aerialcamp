package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/sxc/aerialcamp/models"
)

type Galleries struct {
	Templates struct {
		New Template
	}
	GalleryService *models.GalleryService
}

func (g Galleries) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email     string
		CSRFField template.HTML
	}
	data.Email = r.FormValue("email")
	data.CSRFField = csrf.TemplateField(r)
	g.Templates.New.Execute(w, r, data)
}
