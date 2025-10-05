package main

import (
	"fmt"
	"net/http"

	"groupi-tracker/functions"
)

func main() {
	http.HandleFunc("/", functions.Home)
	http.HandleFunc("/artists/", functions.About)
	http.HandleFunc("/statics/", functions.Style)

	fmt.Println("server started on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to start server: ", err)
		return
	}
}
