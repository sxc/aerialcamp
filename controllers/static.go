package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "What is guanyuchi?",
			Answer:   "Share food photos with your friends.",
		},
		{
			Question: "How to use guanyuchi?",
			Answer:   "Register with email address.",
		},
		{
			Question: "I lost my password, how to reset?",
			Answer:   "Go to the login page and click forgot password.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
