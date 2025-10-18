package functions

import (
	"fmt"
	"net/http"
	"os"
)

// Style serves static files such as CSS or assets requested by the client.
// It checks file validity and permissions before safely sending the file to the response.
func Style(w http.ResponseWriter, r *http.Request) {
	fileinfo, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		RenderError(w, "Please try later", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if fileinfo.IsDir() {
		RenderError(w, "Access denied", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
