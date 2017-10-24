package apimethod

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"log"
)
// 周囲のKokoを取得するGET
type GetKokoAround struct {
	Section       string  `json:"section"`
	LocationLati  float64 `json:"location_lati"` // 緯度
	LocationLongi float64 `json:"location_longi"` // 経度
	OverKokoid    int     `json:"over_kokoid"`
}
// 周囲のKokoを取得するRESPONSE
type KokoAround struct {
	ID            int     `json:"id"`
	LocationLati  float64 `json:"location_lati"`
	LocationLongi float64 `json:"location_longi"`
	Section       string  `json:"section"`
}
type ResponseKokoAround  [] KokoAround

func GetAroundKoko(w http.ResponseWriter, r *http.Request, params httprouter.Params){

	// request body をjsonにエンコード
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var t = GetKokoAround{}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	//response のjsonを作って書き込み
	var responseJson ResponseKokoAround
	responseJson = ResponseKokoAround{
		KokoAround{Section:t.Section},
	}
	json.NewEncoder(w).Encode(responseJson)
	getAroundKokoModel()
	log.Println("access getaroundkoko")
}

func getAroundKokoModel() {

}