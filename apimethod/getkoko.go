package apimethod

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
	"encoding/json"
)
// 指定したkokoを取得する GET
type GetKokos struct {
	kokolist []int `json:"koko"`
}

// kokoオブジェクト
type Koko struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	UserName         string  `json:"user_name"`
	UserIcon         string  `json:"user_icon"`
	Created          time.Time  `json:"created"`
	Content          string  `json:"content"`
	Attach           string  `json:"attach"`
	ChildrenNum      int     `json:"children_children"`
	Like             int     `json:"like"`
	LocationLati     float64 `json:"location-lati"`
	LocationLongi    float64 `json:"location-longi"`
	Section          string  `json:"section"`
	ParentID         int     `json:"parent_id"`
}
// 指定したkokoを取得するRESPONSE
type ResponseKokos[] Koko
func GetKoko(w http.ResponseWriter, params httprouter.Params) {
	getKokoModel()
	response := ResponseKokos{
		Koko{ID:1},
		Koko{ID:2},
	}
	json.NewEncoder(w).Encode(response);

}

func getKokoModel()  {

}