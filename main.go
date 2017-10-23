package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	_"github.com/julienschmidt/httprouter"
	"github.com/julienschmidt/httprouter"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	router := httprouter.New()
	router.GET("/koko:path", koko)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func koko(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, get-koko %q", html.EscapeString(r.URL.Path) )
}
