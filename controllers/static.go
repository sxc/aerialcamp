package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
