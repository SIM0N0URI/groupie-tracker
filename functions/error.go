package functions

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"syscall"
)

// RenderError renders an error page with the specified message and HTTP status code.
// It parses the error template, injects the data, and writes the rendered result to the response writer.
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

	if _, err := buf.WriteTo(w); err != nil {
		if !errors.Is(err, syscall.EPIPE) {
			fmt.Println("Failed to write buffer:", err)
		}
		return
	}
}
