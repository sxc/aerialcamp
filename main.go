package main

import (
	"fmt"
	"net/http"

	// "log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sxc/aerialcamp/controllers"
	"github.com/sxc/aerialcamp/migrations"
	"github.com/sxc/aerialcamp/models"
	"github.com/sxc/aerialcamp/templates"
	"github.com/sxc/aerialcamp/views"

	"github.com/gorilla/csrf"
)

// add executeTemplate
// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse(filepath)
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	userService := models.UserService{
		DB: db,
	}

	sessionService := models.SessionService{
		DB: db,
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database.")
	// us := models.UserService{DB: db}

	// user, err := us.Create("john12@example.com", "password")

	// user, err := us.Create("john01@example.com", "password")

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS,
			"faq.gohtml", "tailwind.gohtml"))))

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.PorcessSignOut)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	// 32 alphabet keys
	csrfKey := "0123456789abcdefsafdsafdsafdsaffS"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false),
	)
	// []byte(csrfKey),
	// TODO: Fix this before deployng
	// csrf.Secure(false))
	// http.ListenAndServe(":3000", csrfMw(r))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", csrfMw(umw.SetUser(r)))
}
