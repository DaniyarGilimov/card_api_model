package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UserRegist single user in registration mode
type UserRegist struct {
	Token         string           `json:"token"`
	PushToken     string           `json:"pushToken" bson:"pushToken"`
	UserID        int              `json:"userId" bson:"userId"`
	Chips         int              `json:"chips"`
	Golds         int              `json:"golds"`
	Username      string           `json:"userName" bson:"userName"`
	Email         string           `json:"email" binding:"required" example:"test@gmail.com"`
	Avatar        string           `json:"photoUrl" bson:"photoUrl"`
	RewardTime    time.Time        `json:"rewardTime" bson:"rewardTime"`
	Inventory     *AllGoods        `json:"inventory" bson:"allGoods"`
	Settings      *Settings        `json:"settings" bson:"settings"`
	OnHat         *StaticInventory `json:"onHat" bson:"onHat"`
	DailyLeft     int              `json:"dailyLeft" bson:"dailyLeft"`
	RoomID        int              `json:"roomId" bson:"roomId"` //gives current playing room
	IsPushEnabled bool             `json:"isPushEnabled" bson:"isPushEnabled"`
	IsBot         bool             `json:"isBot" bson:"isBot"`
	// BoymanData
}

func (s *UserRegist) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&s, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqSettingsUpdate struct {
	Username       string `json:"username" bson:"username"`
	AvatarID       int    `json:"avatarId" bson:"avatarId"`
	HatID          int    `json:"hatId" bson:"hatId"`
	IsPushEnabled  bool   `json:"isPushEnabled" bson:"isPushEnabled"`
	IsSoundEnabled bool   `json:"isSoundEnabled" bson:"isSoundEnabled"`
	Language       string `json:"language" bson:"language"`
}

func (s *ReqSettingsUpdate) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&s, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if s.Username == "" || len(s.Username) > 12 {
		return errors.New("bad request | username is empty")
	}

	return nil
}

type BuyInventory struct {
	ID   int    `json:"id" bson:"id"`
	Type string `json:"type" bson:"type"`
}

func (p *BuyInventory) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID <= 0 {
		return errors.New("bad request | id is invalid")
	}

	if p.Type == "" || (p.Type != "hats" && p.Type != "avatars" && p.Type != "chips" && p.Type != "throwables") {
		return errors.New("bad request | type is empty")
	}

	return nil
}

type ReqPage struct {
	Page int `json:"page" bson:"page"`
}

func (p *ReqPage) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.Page <= 0 {
		return errors.New("bad request | page is invalid")
	}

	return nil
}

type ReqSearch struct {
	Page   int    `json:"page" bson:"page"`
	Search string `json:"search" bson:"search"`
}

func (p *ReqSearch) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.Page <= 0 {
		return errors.New("bad request | page is invalid")
	}

	return nil
}

type ReqSearchOther struct {
	Page   int    `json:"page" bson:"page"`
	Search string `json:"search" bson:"search"`
	ID     int    `json:"id" bson:"id"`
}

func (p *ReqSearchOther) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.Page <= 0 || p.ID <= 0 {
		return errors.New("bad request | page is invalid")
	}

	return nil
}

type ReqEditFriend struct {
	FriendID int    `json:"friendId" bson:"friendId"`
	Type     string `json:"type" bson:"type"`
}

func (p *ReqEditFriend) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.FriendID <= 0 {
		return errors.New("bad request | friendId is invalid")
	}

	if p.Type == "" || (p.Type != "delete" && p.Type != "accept" && p.Type != "request" && p.Type != "reject" && p.Type != "cancel") {
		return errors.New("bad request | type is empty")
	}

	return nil
}

type ReqCheckUsername struct {
	Username string `json:"username" bson:"username"`
}

func (s *ReqCheckUsername) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&s, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if s.Username == "" || len(s.Username) > 12 {
		return errors.New("bad request | username is empty")
	}
	return nil
}

type ReqPush struct {
	ID     int    `json:"id" bson:"id"`
	Title  string `json:"title" bson:"title"`
	Body   string `json:"body" bson:"body"`
	Status string `json:"status" bson:"status"`
}

func (s *ReqPush) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&s, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if s.ID <= 0 {
		return errors.New("bad request | id is invalid")
	}
	return nil
}

type ReqChips struct {
	Chips int64 `json:"chips" bson:"chips"`
}

func (p *ReqChips) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.Chips <= 0 {
		return errors.New("bad request | chips is invalid")
	}

	return nil
}

type ReqLanguage struct {
	Lang string `json:"lang" bson:"lang"`
}

func (p *ReqLanguage) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.Lang == "" {
		return errors.New("bad request | lang is invalid")
	}

	return nil
}

type ReqBet struct {
	Bet int64 `json:"bet" bson:"bet"`
}

func (p *ReqBet) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.Bet <= 0 {
		return errors.New("bad request | bet is invalid")
	}

	return nil
}

type ReqTournament struct {
	TournamentRules []string           `json:"tournamentRules"`
	Prize           []*TournamentPrize `json:"prize"`
	Info            *TournamentInfo    `json:"info"`
	EndTime         time.Time          `json:"endTime"`
	MinLevel        int                `json:"minLevel" bson:"minLevel"`
	PurchaseIAPID   []int              `json:"purchaseIAPID" bson:"purchaseIAPID"`
}

func (p *ReqTournament) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if len(p.TournamentRules) == 0 {
		return errors.New("bad request | tournamentRules is empty")
	}
	if p.Prize == nil {
		return errors.New("bad request | prize is empty")
	}
	if p.Info == nil {
		return errors.New("bad request | info is empty")
	}
	if p.EndTime.IsZero() {
		return errors.New("bad request | endTime is empty")
	}
	return nil
}

type ReqTournamentPurchase struct {
	TournamentID int `json:"tournamentId" bson:"tournamentID"`
	PurchaseID   int `json:"purchaseId" bson:"purchaseID"`
}

func (p *ReqTournamentPurchase) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.TournamentID <= 0 {
		return errors.New("bad request | tournamentId is invalid")
	}
	if p.PurchaseID <= 0 {
		return errors.New("bad request | purchaseId is invalid")
	}
	return nil
}

type AddPurchase struct {
	UserID       int    `json:"userId"`
	Username     string `json:"username"`
	IapIDAndroid string `json:"iapIdAndroid"`
	IapIDIos     string `json:"iapIdIos"`
}

func (p *AddPurchase) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}
	if p.UserID <= 0 && p.Username == "" {
		return errors.New("bad request | userId is invalid")
	}
	if p.IapIDAndroid == "" && p.IapIDIos == "" {
		return errors.New("bad request | iapIdAndroid and iapIdIos is empty")
	}
	return nil
}
