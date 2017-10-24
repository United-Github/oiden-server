package apimethod

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
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

func GetAroundKoko(w http.ResponseWriter, params httprouter.Params){
	kokolist := ResponseKokoAround{
		KokoAround{ID:2},
		KokoAround{ID:1},
	}
	json.NewEncoder(w).Encode(kokolist)
	getAroundKokoModel()
}

func getAroundKokoModel() {

}