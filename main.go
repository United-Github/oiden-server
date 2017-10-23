package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"./config"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(config.Sql_user))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
