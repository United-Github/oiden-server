package main

import "time"

// 周囲のKokoを取得するGET
type GetKokoAround struct {
	Section       string  `json:"section"`
	LocationLati  float64 `json:"location_lati"` // 緯度
	LocationLongi float64 `json:"location_longi"` // 経度
	OverKokoid    int     `json:"over_kokoid"`
}
// 周囲のKokoを取得するRESPONSE
type ResponseKokoAround struct {
	Kokolist []struct {
		ID            int     `json:"id"`
		LocationLati  float64 `json:"location_lati"`
		LocationLongi float64 `json:"location_longi"`
		Section       string  `json:"section"`
	} `json:"kokolist"`
}

// 指定したkokoを取得する GET
type GetKokos struct {
	Koko []int `json:"koko"`
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
type ResponseKokos struct {
	KokoList [] Koko `json:"koko-list"`
}