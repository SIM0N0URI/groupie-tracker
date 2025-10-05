package functions

import (
	"bytes"
	"html/template"
	"net/http"
)

func RenderError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	data := ErrorPage{
		Code:    code,
		Message: message,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, "Template write error", http.StatusInternalServerError)
		return
	}
}
