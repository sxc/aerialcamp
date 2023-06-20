package controllers

import (
	"net/http"

	"github.com/sxc/aerialcamp/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "What is this?",
			Answer:   "This is a sample FAQ question.",
		},
		{
			Question: "Where is this 2?",
			Answer:   "This is a sample FAQ question.",
		},
		{
			Question: "What is this 3?",
			Answer:   "This is a sample FAQ question.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
