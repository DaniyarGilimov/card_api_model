package model

// RequestMessageTelegram is data that telegram sends to us
type RequestMessageTelegram struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
	InlineMessageID string `json:"inline_message_id"`
}

// CallbackMessageTelegram is data that telegram sends to us
type CallbackMessageTelegram struct {
	CallBackQuery struct {
		Message struct {
			MessageID int64  `json:"message_id"`
			Text      string `json:"text"`
			Chat      struct {
				ID int64 `json:"id"`
			} `json:"chat"`
		} `json:"message"`
		Data string `json:"data"`
	} `json:"callback_query"`
}

// SendMessageTelegram is data that we should send to telegram
type SendMessageTelegram struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// TelegramSubscriber in database
type TelegramSubscriber struct {
	ChatID int64 `json:"chat_id" bson:"chat_id"`
}

// DeleteMessageTelegram is data that we should send to telegram
type DeleteMessageTelegram struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}

type ClickByName struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type StuckRoomInfo struct {
	Inform string `json:"inform"`
}

type MetricsDaily struct {
	Day          int `json:"day" bson:"day"`
	Month        int `json:"month" bson:"month"`
	Year         int `json:"year" bson:"year"`
	Shop         int `json:"shop" bson:"shop"`
	DAU          int `json:"dau" bson:"dau"`
	Registration int `json:"registration" bson:"registration"`
	ShowInGame   int `json:"showInGame" bson:"showInGame"`
	BuyInGame    int `json:"buyInGame" bson:"buyInGame"`
}

type TransactionDaily struct {
	Count      int     `json:"count" bson:"count"`
	CountOffer int     `json:"countOffer" bson:"countOffer"`
	CountShop  int     `json:"countShop" bson:"countShop"`
	CheckDay   float64 `json:"checkDay" bson:"checkDay"`
	CheckAvg   float64 `json:"checkAvg" bson:"checkAvg"`
	CheckMonth float64 `json:"checkMonth" bson:"checkMonth"`
}

type DAUInfo struct {
	MetricsDaily
	TransactionDaily

	MAU    int `json:"mau" bson:"mau"`
	CRDAU  int //dau/mau
	CRSHOP int //shop/shop

	// Formated info
	FormatedCheckAvg   string `json:"formatedCheckAvg"`
	FormatedCheckMonth string `json:"formatedCheckMonth"`
}
