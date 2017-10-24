package main

import (
	"./apimethod"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*ココの個別データを取得するメソッド*/
func koko(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setContentHeader(w)
	apimethod.GetKoko(w, params)
}

/* デバイスが存在する周囲のココをリストアップするメソッド */
func kokoaround(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setContentHeader(w)
	apimethod.GetAroundKoko(w, params)
}
func setContentHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.GET("/koko", koko)
	router.GET("/kokoaround", kokoaround)
	router.ServeFiles("/image/*filepath", http.Dir("image/"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
