package utils

import (
	"html/template"
	"net/http"
)

// RenderTemplate parses and executes the provided page templates and writes the
// result to the http.ResponseWriter. Intended for any server-rendered pages.
func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("web/template/base.html", "web/template/sidebar.html", "web/template/"+page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

// AuthTemplate renders authentication-specific templates (login/register).
func AuthTemplate(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("web/template/authentication/base.html", "web/template/authentication/"+page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}
