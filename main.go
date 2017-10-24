package main

import (
	"./apimethod"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*ココの個別データを取得するメソッド*/
func koko(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setContentJsonHeader(w)
	apimethod.GetKoko(w, r, params)
}

/* デバイスが存在する周囲のココをリストアップするメソッド */
func kokoaround(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setContentJsonHeader(w)
	apimethod.GetAroundKoko(w, r, params)
}

/* ココをポストする */
func postkoko(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setContentJsonHeader(w)
	apimethod.PostKoko(w, r, params)
}

/* APIのHTTP-HeaderにJsonを入れる*/
func setContentJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.POST("/koko", koko)
	router.POST("/kokoaround", kokoaround)
	router.POST("/postkoko", postkoko)
	router.ServeFiles("/image/*filepath", http.Dir("image/"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
