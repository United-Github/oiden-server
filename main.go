package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func koko(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "koko")
}
func kokoaround(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "kokoaround")
}

func main() {
	router := httprouter.New()
	router.GET("/koko", koko)
	router.GET("/kokoaround", kokoaround)
	router.ServeFiles("/image/*filepath", http.Dir("image/"))
	log.Fatal(http.ListenAndServe(":8080", router))
}


