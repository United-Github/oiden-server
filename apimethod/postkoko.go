package apimethod

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"log"
)
// 周囲のKokoを取得するGET
type postKokoRequest struct {
	UserID        int    `json:"user_id"`
	Content       string `json:"content"`
	Attach        string `json:"attach"`
	LocationLati  int    `json:"location-lati"`
	LocationLongi int    `json:"location-longi"`
}
// 周囲のKokoを取得するRESPONSE
type postKokoResponse struct {
	Result int `json:"result"`
}

func PostKoko(w http.ResponseWriter, r *http.Request, params httprouter.Params){

	// request body をjsonにエンコード
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var t = postKokoRequest{}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	//response のjsonを作って書き込み
	var responseJson postKokoResponse
	responseJson = postKokoResponse{Result:200}
	json.NewEncoder(w).Encode(responseJson)
	getAroundKokoModel()
	log.Println("access getaroundkoko")
}

func getPostKokoModel() {

}