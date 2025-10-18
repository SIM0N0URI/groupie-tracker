package functions

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"syscall"
)

var Artists []Artist

// Home handles the main page request and displays the list of all artists.
// It validates the request, fetches artist data from the API, and renders the index template.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to parse index template", err)
		return
	}

	var buf bytes.Buffer

	err = FetchJson(url, &Artists)
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err = tmpl.Execute(&buf, &Artists)
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println("Failed to execute index template", err)
		return
	}

	if _, err := buf.WriteTo(w); err != nil {
		if !errors.Is(err, syscall.EPIPE) {
			fmt.Println("Failed to write buffer:", err)
		}
		return
	}
}
