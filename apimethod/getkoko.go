package apimethod

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
	"encoding/json"
	"log"
)
// 指定したkokoを取得する GET
type GetKokoList[] int

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
func GetKoko(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var responseJson ResponseKokos
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var list GetKokoList
	err := decoder.Decode(&list)
	if err != nil {
		panic(err)
	}
	for _, value := range list {
		t := Koko{ID: value}
		responseJson = append(responseJson, t)
	}
	json.NewEncoder(w).Encode(responseJson)
	log.Println("access getkoko")
}

func getKokoModel()  {

}