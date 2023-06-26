package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	tpl := template.New(pattern[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented")
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %v", err)
	}

	return Template{HTMLTpl: tpl}, nil
}

type Template struct {
	HTMLTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.HTMLTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error cloning the template.", http.StatusInternalServerError)
		return
	}
	tpl = tpl.Funcs(
		template.FuncMap{
			// "csrfField": func() template.HTML {
			// 	return csrf.TemplateField(r)
			// },
		},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
