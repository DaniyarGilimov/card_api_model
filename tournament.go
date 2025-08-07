package model

import "time"

type TournamentDisplay struct {
	Tournament
	Shop      []*TournamentShop `json:"shop" bson:"shop"`
	Current   *PersonTournament `json:"current" bson:"current"`
	IsLocked  bool              `json:"isLocked" bson:"isLocked"`
	WhyLocked string            `json:"whyLocked" bson:"whyLocked"`
}

type TournamentShop struct {
	PurchaseID      int                  `json:"purchaseId" bson:"purchaseID"`
	TournamentChips int64                `json:"tournamentChips" bson:"tournamentChips"`
	Image           string               `json:"image" bson:"image"`
	Purchase        *PurchaseByRealMoney `json:"purchase" bson:"purchase"`
	PurchaseChips   *PurchaseByChips     `json:"purchaseChip" bson:"purchaseChips"`
	PurchaseGolds   *PurchaseByGolds     `json:"purchaseGold" bson:"purchaseGolds"`
	Type            string               `json:"type" bson:"type"`
}

type PersonTournament struct {
	TournamentChips int64   `json:"tournamentChips" bson:"tournamentChips"`
	Place           int64   `json:"place" bson:"place"`
	Percentage      float64 `json:"percentage" bson:"percentage"`
}

type Tournament struct {
	TournamentID    int                `json:"id" bson:"tournamentID"`
	TournamentRules []string           `json:"tournamentRules" bson:"tournamentRules"`
	Prize           []*TournamentPrize `json:"prize" bson:"prize"`
	Info            *TournamentInfo    `json:"info" bson:"info"`
	EndTime         time.Time          `json:"endTime" bson:"endTime"`
	Status          string             `json:"status" bson:"status"`
	MinLevel        int                `json:"minLevel" bson:"minLevel"`
	PurchaseIAPID   []int              `json:"purchaseIAPID" bson:"purchaseIAPID"`
}

type UserTournament struct {
	UserID          int       `json:"userID" bson:"userID"`
	TournamentID    int       `json:"tournamentID" bson:"tournamentID"`
	TournamentChips int64     `json:"tournamentChips" bson:"tournamentChips"`
	LastUpdated     time.Time `json:"lastUpdated" bson:"lastUpdated"`
}

type TournamentPrize struct {
	StartPlace int   `json:"startPlace" bson:"startPlace"`
	EndPlace   int   `json:"endPlace" bson:"endPlace"`
	Chips      int64 `json:"chips" bson:"chips"`
	Golds      int64 `json:"golds" bson:"golds"`
}

type TournamentInfo struct {
	Image   string `json:"image" bson:"image"`
	Title   string `json:"title" bson:"title"`
	SecLeft int64  `json:"secLeft" bson:"secLeft"`
	Chips   int64  `json:"chips" bson:"chips"`
	Golds   int64  `json:"golds" bson:"golds"`
}

type PurchaseByRealMoney struct {
	IapIDIos     string `json:"iapIdIos" bson:"iapIdIos"`
	IapIDAndroid string `json:"iapIdAndroid" bson:"iapIdAndroid"`
}

type PurchaseByChips struct {
	PurchaseID    int   `json:"purchaseId" bson:"purchaseID"`
	PurchaseChips int64 `json:"purchaseChip" bson:"purchaseChips"`
}

type PurchaseByGolds struct {
	PurchaseID    int   `json:"purchaseId" bson:"purchaseID"`
	PurchaseGolds int64 `json:"purchaseGold" bson:"purchaseGolds"`
}

type TournamentFriend struct {
	Friend          *FriendOutput `json:"friend" bson:"friend"`
	TournamentChips int64         `json:"tournamentChips" bson:"tournamentChips"`
	Golds           int64         `json:"golds" bson:"golds"`
	Chips           int64         `json:"chips" bson:"chips"`
}

type TournamentLeaders struct {
	Leaders []*TournamentFriend `json:"leaders" bson:"leaders"`
	Winners []*TournamentFriend `json:"winners" bson:"winners"`
}
