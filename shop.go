package model

import "time"

// AllGoods describes set of all goods
type AllGoods struct {
	Throwables           []*ThrowableInventory `json:"throwables" bson:"throwables"`
	Hats                 []*StaticInventory    `json:"hats" bson:"hats"`
	Avatars              []*StaticInventory    `json:"avatars" bson:"avatars"`
	Golds                []*GoldModel          `json:"golds" bson:"golds"`
	ChipsMoney           []*GoldModel          `json:"chipsMoney" bson:"chipsMoney"`
	TournamentChipsMoney []*GoldModel          `json:"tournamentChipsMoney" bson:"tournamentChipsMoney"`
	Chips                []*ChipModel          `json:"chips" bson:"chips"`
	SpecialOffers        []*SpecialOffer       `json:"specialOffers" bson:"specialOffers"`
	TournamentShops      []*TournamentShop     `json:"tournamentShops" bson:"tournamentShops"`
}

type SpecialOffer struct {
	OfferID      string      `json:"offerId" bson:"offerId"`
	ChipsIndex   int         `json:"chipsIndex" bson:"chipsIndex"`
	RankIndex    int         `json:"rankIndex" bson:"rankIndex"`
	ShortTitle   string      `json:"shortTitle" bson:"shortTitle"`
	Title        string      `json:"title" bson:"title"`
	PriceModel   *PriceModel `json:"priceModel" bson:"priceModel"`
	IapIDAndroid string      `json:"iapIdAndroid" bson:"iapIdAndroid"`
	IapIDIos     string      `json:"iapIdIos" bson:"iapIdIos"`
	Image        string      `json:"image" bson:"image"`
	SecLeft      int         `json:"secLeft" bson:"-"`

	Golds     int64 `json:"golds" bson:"golds"`
	Chips     int64 `json:"chips" bson:"chips"`
	OldAmount int64 `json:"oldAmount" bson:"oldAmount"`

	Throwables []*ThrowableInventory `json:"throwables" bson:"throwables"`
	Hats       []*StaticInventory    `json:"hats" bson:"hats"`
	Avatars    []*StaticInventory    `json:"avatars" bson:"avatars"`
}

type ChipModel struct {
	ID         int         `json:"id" bson:"id"`
	Title      string      `json:"title" bson:"title"`
	Amount     int64       `json:"amount" bson:"amount"`
	PriceModel *PriceModel `json:"priceModel" bson:"priceModel"`
}

type PriceModel struct {
	Price          int     `json:"price" bson:"price"`
	DiscountPrice  float32 `json:"discountPrice" bson:"discountPrice"`
	IsPriceInChips bool    `json:"isPriceInChips" bson:"isPriceInChips"`
}

type ChipsWithMoneyModel struct {
	ID           int         `json:"id" bson:"id"`
	Title        string      `json:"title" bson:"title"`
	Amount       int         `json:"amount" bson:"amount"`
	PriceModel   *PriceModel `json:"priceModel" bson:"priceModel"`
	IapIDAndroid string      `json:"iapIdAndroid" bson:"iapIdAndroid"`
	IapIDIos     string      `json:"iapIdIos" bson:"iapIdIos"`
}

type GoldModel struct {
	ID           int         `json:"id" bson:"id"`
	Title        string      `json:"title" bson:"title"`
	Amount       int64       `json:"amount" bson:"amount"`
	PriceModel   *PriceModel `json:"priceModel" bson:"priceModel"`
	IapIDAndroid string      `json:"iapIdAndroid" bson:"iapIdAndroid"`
	IapIDIos     string      `json:"iapIdIos" bson:"iapIdIos"`

	// used only in chips
	OldChips        int64   `json:"oldChips"`
	BuyBonusPercent float64 `json:"buyBonusPercent" bson:"buyBonusPercent"`
}

// ThrowableInventory goods wit price range
type ThrowableInventory struct {
	ID          int         `json:"id" bson:"id"`
	Title       string      `json:"title" bson:"title"`
	Amount      int         `json:"amount" bson:"amount"`
	Description string      `json:"description" bson:"description"`
	ThrowID     int         `json:"throwId" bson:"throwId"`
	PriceModel  *PriceModel `json:"priceModel" bson:"priceModel"`
}

// TemporaryInventory  goods with time limit in days temporaries
type TemporaryInventory struct {
	ID               int                   `json:"id" bson:"id"`
	Title            string                `json:"title" bson:"title"`
	TitleTuple       *Tuple                `json:"-" bson:"titleTuple"`
	Description      string                `json:"description" bson:"description"`
	DescriptionTuple *Tuple                `json:"-" bson:"descriptionTuple"`
	PriceRange       []*DurationPriceRange `json:"priceRange" bson:"priceRange"`
	Duration         int                   `json:"duration" bson:"duration"`
}

// StaticInventory goods that is bought once, and used forever
type StaticInventory struct {
	ID          int         `json:"id" bson:"id"`
	Title       string      `json:"title" bson:"title"`
	Description string      `json:"description" bson:"description"`
	IsPurchased bool        `json:"isPurchased" bson:"isPurchased"`
	Image       string      `json:"image" bson:"image"`
	Rarity      float32     `json:"rarity" bson:"rarity"`
	PriceModel  *PriceModel `json:"priceModel" bson:"priceModel"`
}

// PriceRange used to describe price range
type PriceRange struct {
	ID     int `json:"id" bson:"id"`
	Amount int `json:"amount" bson:"amount"`
	Price  int `json:"price" bson:"price"`
}

// DurationPriceRange used to describe price range in duration
type DurationPriceRange struct {
	ID       int `json:"id" bson:"id"`
	Duration int `json:"duration" bson:"duration"`
	Price    int `json:"price" bson:"price"`
}

// ShopChip is chip data
type ShopChip struct {
	ID             int    `json:"id"`
	ChipName       string `json:"chipName" bson:"chipName"`
	ChipNameTuple  *Tuple `json:"-" bson:"chipNameTuple"`
	ChipCount      int    `json:"chipCount" bson:"chipCount"`
	Price          int    `json:"price"`
	RealCount      int    `json:"realCount" bson:"realCount"`
	SalePercentage int    `json:"salePercentage" bson:"salePercentage"`
}

// Tuple for languages
type Tuple struct {
	RU string `json:"RU" bson:"RU"`
	EN string `json:"EN" bson:"EN"`
	HI string `json:"HI" bson:"HI"`
}

type Transaction struct {
	ID             string                   `json:"id" bson:"id"`
	ProductID      string                   `json:"productId" bson:"productId"`
	OrderID        string                   `json:"orderId" bson:"orderId"`
	Time           string                   `json:"time" bson:"time"`
	FactTime       time.Time                `json:"factTime" bson:"factTime"`
	GoogleResponse *ReadGoogleValidationIAP `json:"googleResponse" bson:"googleResponse"`
	AppleResponse  *ReadAppleValidation     `json:"appleResponse" bson:"appleResponse"`
	UserID         int                      `json:"userId" bson:"userId"`
}

type DailyRewardPage struct {
	IsAvailable bool         `json:"isAvailable"`
	Rewards     []*DayReward `json:"rewards"`
}

type DayReward struct {
	Day               int   `json:"day" bson:"day"`
	Chips             int64 `json:"chips" bson:"chips"`
	OldChips          int64 `json:"oldChips" bson:"oldChips"`
	DailyBonusPercent int   `json:"dailyBonusPercent" bson:"dailyBonusPercent"`

	Eggs    int    `json:"eggs" bson:"eggs"`
	Tomatos int    `json:"tomatos" bson:"tomatos"`
	Status  string `json:"status" bson:"status"`
}

type IAPData struct {
	Key          string `json:"key" bson:"key"`
	IapIdAndroid string `json:"iapIdAndroid" bson:"iapIdAndroid"`
	IapIdIos     string `json:"iapIdIos" bson:"iapIdIos"`
}

type OfferPage struct {
	IsAvailable  bool `json:"isAvailable" bson:"isAvailable"`
	IsShowDialog bool `json:"isShowDialog" bson:"isShowDialog"`

	Offer *SpecialOffer `json:"offer" bson:"offer"`
}

type Offer struct {
	OfferID      string                `json:"offerId" bson:"offerId"`
	ShortTitle   string                `json:"shortTitle" bson:"shortTitle"`
	Title        string                `json:"title" bson:"title"`
	IAPIDIOS     string                `json:"iapIdIos" bson:"iapIdIos"`
	IAPIDAndroid string                `json:"iapIdAndroid" bson:"iapIdAndroid"`
	SecLeft      int                   `json:"secLeft" bson:"-"`
	Chips        int                   `json:"chips" bson:"chips"`
	Image        string                `json:"image" bson:"image"`
	Throwables   []*ThrowableInventory `json:"throwables" bson:"throwables"`
	Hats         []*StaticInventory    `json:"hats" bson:"hats"`
	Avatars      []*StaticInventory    `json:"avatars" bson:"avatars"`
}

type BarabanIndex struct {
	Index              int
	BarabanIndexStart  int
	BarabanIndexEnd    int
	TripleCoefficiency float64
	DoubleCoefficiency float64
}
