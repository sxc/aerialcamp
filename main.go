package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sxc/aerialcamp/views"
	"github.com/sxc/aerialcamp/controllers"
	"github.com/sxc/aerialcamp/templates"
)



// add executeTemplate
func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}



// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	tplPath := filepath.Join("templates", "contact.gohtml")
// 	executeTemplate(w, tplPath)
// 	}

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	tplPath := filepath.Join("templates", "faq.gohtml")
// 	executeTemplate(w, tplPath)
// 	}



func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
