package model

type BoymanData struct {
	Boyman    string `json:"boyman" bson:"boyman" binding:"required" example:"test"`
	Timestamp string `json:"timestamp" bson:"timestamp" binding:"required" example:"test"`
}

// Languages struct
type Languages struct {
	En string `json:"en" bson:"en"`
	Es string `json:"es" bson:"es"`
	De string `json:"de" bson:"de"`
	Fr string `json:"fr" bson:"fr"`
	It string `json:"it" bson:"it"`
	Pt string `json:"pt" bson:"pt"`
	Ru string `json:"ru" bson:"ru"`
	Hi string `json:"hi" bson:"hi"`
}
