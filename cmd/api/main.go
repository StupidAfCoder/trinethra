package main

import (
	"net/http"
	"trinethra/internal/handler"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/analyze", handler.Analyze)

	http.ListenAndServe(":8080", nil)
}
