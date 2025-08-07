package model

// DefaultResponse used in every response
type DefaultResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"result"`
}

// SuccessResponse returns success
func SuccessResponse() *DefaultResponse {
	dr := &DefaultResponse{StatusCode: 1000, Message: "Success"}
	return dr
}

type DefaultPage struct {
	TotalPages  int64       `json:"totalPages" bson:"totalPages"`
	CurrentPage int         `json:"currentPage" bson:"currentPage"`
	Data        interface{} `json:"data" bson:"data"`
}

type FriendPage struct {
	TotalPages     int64       `json:"totalPages" bson:"totalPages"`
	CurrentPage    int         `json:"currentPage" bson:"currentPage"`
	CountOfFriends int         `json:"countOfFriends" bson:"countOfFriends"`
	Data           interface{} `json:"data" bson:"data"`
}

type IsSuccess struct {
	IsSuccess bool `json:"isSuccess" bson:"isSuccess"`
}
