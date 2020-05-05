package main

import (
	"fmt"
	"net/http"
	app "net/http"
)

func handler(w app.ResponseWriter, r *http.Request) {
	app.ServeFile(w, r, "./static/index.html")
}

func main() {
	app.HandleFunc("/", handler)
	fmt.Println("Server running ...")
	app.ListenAndServe(":8080", nil)
}
