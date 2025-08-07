package model

type FriendOutput struct {
	ID               int              `json:"userId" bson:"userId"`
	Status           string           `json:"status" bson:"status"`
	Username         string           `json:"username" bson:"username"`
	Inventory        *PersonInventory `json:"inventory" bson:"inventory"`
	CommonFriends    []string         `json:"commonFriends" bson:"commonFriends"`
	CommonFriendsNum int              `json:"commonFriendsNum" bson:"commonFriendsNum"`
	ChipsWon         int              `json:"chipsWon" bson:"chipsWon"`
}
