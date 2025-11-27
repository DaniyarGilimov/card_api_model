package model

type RoomBet struct {
	Bet []int `json:"bet" bson:"bet"`
}

type RoomOptions struct {
	RoomOptions []*RoomOption `json:"roomOptions" bson:"roomOptions"`
}

type RoomOption struct {
	Bet         int64   `json:"bet" bson:"bet"`
	ChipsAmount []int64 `json:"chipsAmount" bson:"chipsAmount"`
	PlaceName   string  `json:"placeName" bson:"placeName"`
	Wallpaper   string  `json:"wallpaper" bson:"wallpaper"`
}

type SingleMessage struct {
	Message         string    `json:"message"`
	MessageLanguage *Langauge `json:"-"`
}

type SingleMessageWithLang struct {
	MessageID       int       `json:"messageId"`
	Message         string    `json:"message"`
	MessageLanguage *Langauge `json:"messageLanguage"`
}

type Langauge struct {
	En string `json:"en"`
	Ru string `json:"ru"`
	Es string `json:"es"`
	Pr string `json:"pr"`

	It string `json:"it"`
	Hi string `json:"hi"`
}

type ClosedRoomUserActivity struct {
	UserID   int    `json:"userId" bson:"userId"`
	UserName string `json:"userName" bson:"userName"`
	Day      int    `json:"day" bson:"day"`
	ChipsWon int64  `json:"chipsWon" bson:"chipsWon"`
}
