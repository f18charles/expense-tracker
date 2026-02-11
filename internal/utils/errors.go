package utils

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Code    int
	Message string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, msg string) {
	tmpl, err := template.ParseFiles("/web/template/base.html", "/web/template/error_page.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)

	data := ErrorData{
		Code:    code,
		Message: msg,
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
