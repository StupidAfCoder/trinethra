package main

import (
	"fmt"
	"net/http"
	"trinethra/internal/handler"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/analyze", handler.Analyze)

	fmt.Print("Server Is Live!!")
	http.ListenAndServe(":8080", nil)
}
