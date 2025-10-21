package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

// User single user
type User struct {
	Token     string `json:"token"`
	PushToken string `json:"pushToken" bson:"pushToken"`
	UserID    int    `json:"userId" bson:"userId"`
	Username  string `json:"username" bson:"username"`
	Email     string `json:"email"`

	RewardData   *RewardData      `json:"rewardData" bson:"rewardData"`
	Inventory    *PersonInventory `json:"inventory" bson:"inventory"`
	Settings     *Settings        `json:"settings" bson:"settings"`
	OnHat        *StaticInventory `json:"onHat" bson:"onHat"`
	TimeData     *TimeData        `json:"timeData" bson:"timeData"`
	OfferData    *OfferData       `json:"offerData" bson:"offerData"`
	OfferHistory *OfferHistory    `json:"offerHistory" bson:"offerHistory"`
	SpinHistory  SpinHistory      `json:"spinHistory" bson:"spinHistory"`
	Tournament   *UserTournament  `json:"tournament" bson:"tournament"`

	RoomID       int       `json:"roomId" bson:"roomId"` //gives current playing room
	AdvertTaken  int       `json:"advertTaken" bson:"advertTaken"`
	Achievements []int     `json:"achievements" bson:"achievements"`
	JoinData     *JoinData `json:"-" bson:"joinData,omitempty"`
	Deleted      bool      `json:"deleted" bson:"deleted"`
	IsBot        bool      `json:"isBot" bson:"isBot"`
	Warnings     int       `json:"warnings" bson:"warnings"`

	// api v3 fileds
	ImageUrl         string `json:"imageUrl" bson:"imageUrl"`
	Chips            int64  `json:"chips" bson:"chips"`
	RateClaimDone    bool   `json:"rateClaimDone" bson:"rateClaimDone"`
	PlinkoCount      int    `json:"plinkoCount" bson:"plinkoCount"`
	MysteryCardCount int    `json:"mysteryCardCount" bson:"mysteryCardCount"`
}

func (u *User) AddV3Fields() {
	u.Chips = u.Inventory.Chips
	u.ImageUrl = u.Inventory.Avatars[0].Image // Assuming the first hat is the active one
}

type SpinHistory struct {
	TotalSpinDone int         `json:"totalSpinDone" bson:"totalSpinDone"`
	WinDelta      int64       `json:"winDelta" bson:"winDelta"`
	SpinBuyInfo   SpinBuyInfo `json:"spinBuyInfo" bson:"spinBuyInfo"`
}

type OfferHistory struct {
	LastOfferEnd   time.Time `json:"lastOfferEnd" bson:"lastOfferEnd"`
	PushDone       bool      `json:"pushDone" bson:"pushDone"`
	PurchasedOffer bool      `json:"purchasedOffer" bson:"purchasedOffer"`
	PurchasedChips bool      `json:"purchasedChips" bson:"purchasedChips"`
}
type RewardData struct {
	PushDone bool `json:"pushDone" bson:"pushDone"`
	Streak   int  `json:"streak" bson:"streak"`
}

type TimeData struct {
	Registration       time.Time `json:"registration" bson:"registration"`
	DailyClaim         time.Time `json:"dailyClaim" bson:"dailyClaim"` // will be deleted, but need for old version
	DailyClaimDeadLine time.Time `json:"dailyClaimDeadLine" bson:"dailyClaimDeadLine"`
	LastDAUAdd         time.Time `json:"lastDAUAdd" bson:"lastDAUAdd"`
	LastShopAdd        time.Time `json:"lastShopAdd" bson:"lastShopAdd"`
	LastSpin           time.Time `json:"lastSpin" bson:"lastSpin"`
	LastPlinko         time.Time `json:"lastPlinko" bson:"lastPlinko"`
	LastMysteryCard    time.Time `json:"lastMysteryCard" bson:"lastMysteryCard"`
}
type Settings struct {
	Version        int    `json:"version"`
	IsPushEnabled  bool   `json:"isPushEnabled" bson:"isPushEnabled"`
	IsSoundEnabled bool   `json:"isSoundEnabled" bson:"isSoundEnabled"`
	Language       string `json:"language" bson:"language"`
}

type PersonInventory struct {
	PlayerID int   `json:"playerId" bson:"playerId"`
	ChipsWon int64 `json:"chipsWon" bson:"chipsWon"`
	Golds    int64 `json:"golds" bson:"golds"`
	Chips    int64 `json:"chips" bson:"chips"`

	ChipsDeltaTotal int64 `json:"-" bson:"-"`

	TodayChips int64 `json:"todayChips" bson:"todayChips"`
	//TotalChips      int64                 `json:"totalChips" bson:"-"` // used in game total amount of chips, variable chips above will be used as handChips
	ActiveHatID    int                   `json:"activeHatId" bson:"activeHatId"`
	ActiveAvatarID int                   `json:"activeAvatarId" bson:"activeAvatarId"`
	Throwables     []*ThrowableInventory `json:"throwables" bson:"throwables"`
	Hats           []*StaticInventory    `json:"hats" bson:"hats"`
	Avatars        []*StaticInventory    `json:"avatars" bson:"avatars"`
	Level          PersonLevel           `json:"level" bson:"level"`
}

func (pi *PersonInventory) Duplicate(TotalBetAmount int64) *PersonInventory {
	copy := *pi
	copy.Chips -= TotalBetAmount
	return &copy
}

type PersonLevel struct {
	LevelInfo
	CurrentLevel int `json:"currentLevel" bson:"currentLevel"`
	CurrentXP    int `json:"currentXp" bson:"currentXp"`
	TargetXP     int `json:"targetXp" bson:"targetXp"`
}

type LevelInfo struct {
	ID                int     `json:"id" bson:"id"`
	RankIndex         int     `json:"rankIndex" bson:"rankIndex"`
	Name              string  `json:"name" bson:"name"`
	Color             string  `json:"color" bson:"color"`
	MinLevel          int     `json:"minLevel" bson:"minLevel"`
	MaxLevel          int     `json:"maxLevel" bson:"maxLevel"`
	BuyBonusPercent   float64 `json:"buyBonusPercent" bson:"buyBonusPercent"`
	DailyBonusPercent float64 `json:"dailyBonusPercent" bson:"dailyBonusPercent"`
	LevelXP           int     `json:"-" bson:"levelXP"`
}

type JoinData struct {
	FriendList *Friends `json:"friends" bson:"friends"`
}

// Friends describes all friends
type Friends struct {
	UserID      int   `json:"userId"  bson:"userId"`
	Friends     []int `json:"friends"`
	OutRequests []int `json:"outRequests" bson:"outRequests"`
	InRequests  []int `json:"inRequests" bson:"inRequests"`
}

// FriendsPub describes all friends
type FriendsPub struct {
	Friends    []*OtherUser `json:"friends"`
	InRequests []*OtherUser `json:"inRequests"`
}

// Statistics describes statistics
type Statistics struct {
	UserID     int `json:"userId" bson:"userId"`
	Wins       int `json:"wins"`
	Looses     int `json:"looses"`
	InGames    int `json:"inGames" bson:"inGames"`
	Folds      int `json:"folds"`
	TotalGames int `json:"totalGames" bson:"totalGames"`
	Taken      int `json:"taken" bson:"taken"`
	Thrown     int `json:"thrown"`
}

// OtherUser describes other user
type OtherUser struct {
	PushToken        string           `json:"-" bson:"pushToken"`
	UserID           int              `json:"userId" bson:"userId"`
	Username         string           `json:"userName" bson:"userName"`
	DisplayName      string           `json:"displayName" bson:"displayName"`
	Email            string           `json:"email"`
	Avatar           string           `json:"photoUrl" bson:"photoUrl"`
	Chips            int              `json:"chips"`
	RegistrationTime time.Time        `json:"registrationTime" bson:"registrationTime"`
	IsFriend         string           `json:"isFriend" bson:"isFriend"`
	OnHat            *StaticInventory `json:"onHat" bson:"onHat"`
	RoomID           int              `json:"roomId" bson:"roomId"` //gives current playing room
	IsPushEnabled    bool             `json:"-" bson:"isPushEnabled"`
	Language         string           `json:"language"`
	Deleted          bool             `json:"deleted" bson:"deleted"`
}

// OwnData ads
type OwnData struct {
	PersonInventory
	Statistics   *Statistics `json:"-"`
	Achievements []int       `json:"achievements" bson:"achievements"`
}

type Notification struct {
	ID          string                 `json:"id" bson:"id"`
	Title       string                 `json:"title" bson:"title"`
	Image       string                 `json:"image" bson:"image"`
	IsRead      bool                   `json:"isRead" bson:"isRead"`
	Description string                 `json:"description" bson:"description"`
	Action      map[string]interface{} `json:"action" bson:"action"`
}

type ReqNotification struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Image       string    `json:"image" bson:"image"`
	IsRead      bool      `json:"isRead" bson:"isRead"`
	Description string    `json:"description" bson:"description"`
	Action      ReqAction `json:"action" bson:"action"`
}
type ReqAction struct {
	Status string  `json:"status" bson:"status"`
	Data   ReqData `json:"data" bson:"data"`
}
type ReqData struct {
	TakerID    int    `json:"takerId" bson:"takerId"`
	AndroidUrl string `json:"androidUrl" bson:"androidUrl"`
	IosUrl     string `json:"iosUrl" bson:"iosUrl"`
	ButtonText string `json:"buttonText" bson:"buttonText"`
}

func (p *ReqNotification) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(p); err != nil {
		return err
	}
	return nil
}

type NotificationOutput struct {
	ID          string                 `json:"id" bson:"_id"`
	SenderID    int                    `json:"senderId" bson:"senderId"`
	TakerID     int                    `json:"takerId" bson:"takerId"`
	Username    string                 `json:"username" bson:"username"`
	Title       string                 `json:"title" bson:"title"`
	Image       string                 `json:"image" bson:"image"`
	IsRead      bool                   `json:"isRead" bson:"isRead"`
	Description string                 `json:"description" bson:"description"`
	Action      map[string]interface{} `json:"action" bson:"action"`
}

type NotificationAction struct {
	Status string                 `json:"status" bson:"status"`
	Data   map[string]interface{} `json:"data" bson:"data"`
}

type Subscription struct {
	Subscribed             bool   `json:"subscribed" bson:"subscribed"`
	ProductID              string `json:"productID" bson:"productID"`
	ServerVerificationData string `json:"serverVerificationData" bson:"serverVerificationData"`
	ExpireDate             int64  `json:"expireDate" bson:"expireDate"`
	Source                 string `json:"source" bson:"source"`
}

type ReadGoogleValidation struct {
	ExpireTime           string `json:"expiryTimeMillis"`
	StartTime            string `json:"startTimeMillis"`
	AutoRenewing         bool   `json:"autoRenewing"`
	PriceCurrencyCode    string `json:"priceCurrencyCode"`
	PriceAmountMicros    string `json:"priceAmountMicros"`
	CountryCode          string `json:"countryCode"`
	DeveloperPayload     string `json:"developerPayload"`
	PaymentState         int    `json:"paymentState"`
	OrderId              string `json:"orderId"`
	LinkedPurchaseToken  string `json:"linkedPurchaseToken"`
	AcknowledgementState int    `json:"acknowledgementState"`
	Kind                 string `json:"kind"`
}
type ReadGoogleValidationIAP struct {
	PurchaseTimeMillis   string `json:"purchaseTimeMillis"`
	PurchaseState        int    `json:"purchaseState"`
	ConsumptionState     int    `json:"consumptionState"`
	DeveloperPayload     string `json:"developerPayload"`
	OrderID              string `json:"orderId"`
	AcknowledgementState int    `json:"acknowledgementState"`
	Kind                 string `json:"kind"`
	RegionCode           string `json:"regionCode"`
}
type ReadAppleValidation struct {
	Receipt struct {
		ReceiptType                string `json:"receipt_type"`
		AdamID                     int    `json:"adam_id"`
		AppItemID                  int    `json:"app_item_id"`
		BundleID                   string `json:"bundle_id"`
		ApplicationVersion         string `json:"application_version"`
		DownloadID                 int64  `json:"download_id"`
		VersionExternalIdentifier  int    `json:"version_external_identifier"`
		ReceiptCreationDate        string `json:"receipt_creation_date"`
		ReceiptCreationDateMs      string `json:"receipt_creation_date_ms"`
		ReceiptCreationDatePst     string `json:"receipt_creation_date_pst"`
		RequestDate                string `json:"request_date"`
		RequestDateMs              string `json:"request_date_ms"`
		RequestDatePst             string `json:"request_date_pst"`
		OriginalPurchaseDate       string `json:"original_purchase_date"`
		OriginalPurchaseDateMs     string `json:"original_purchase_date_ms"`
		OriginalPurchaseDatePst    string `json:"original_purchase_date_pst"`
		OriginalApplicationVersion string `json:"original_application_version"`
		InApp                      []struct {
			Quantity                string `json:"quantity"`
			ProductID               string `json:"product_id"`
			TransactionID           string `json:"transaction_id"`
			OriginalTransactionID   string `json:"original_transaction_id"`
			PurchaseDate            string `json:"purchase_date"`
			PurchaseDateMs          string `json:"purchase_date_ms"`
			PurchaseDatePst         string `json:"purchase_date_pst"`
			OriginalPurchaseDate    string `json:"original_purchase_date"`
			OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
			OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
			IsTrialPeriod           string `json:"is_trial_period"`
			InAppOwnershipType      string `json:"in_app_ownership_type"`
		} `json:"in_app"`
	} `json:"receipt"`
	Environment   string `json:"environment"`
	LatestReceipt string `json:"latest_receipt"`
	Status        int    `json:"status"`
}

type Validate struct {
	PurchaseID             string `json:"purchaseID" bson:"purchaseID"`
	ProductID              string `json:"productID" bson:"productID"`
	LocalVerificationData  string `json:"localVerificationData" bson:"localVerificationData"`
	ServerVerificationData string `json:"serverVerificationData" bson:"serverVerificationData"`
	Source                 string `json:"source" bson:"source"`
	PackageName            string `json:"packageName" bson:"packageName"`
}

type Achievement struct {
	ID          int     `json:"id" bson:"id"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Reached     bool    `json:"reached" bson:"reached"`
	HavePercent float32 `json:"havePercent" bson:"havePercent"`
	Frequency   string  `json:"frequency" bson:"frequency"`
	AimedFor    string  `json:"aimedFor" bson:"aimedFor"`
	Image       string  `json:"image" bson:"image"`
}

type AppStarted struct {
	MinVersionIOS       int   `json:"minVersionIOS" bson:"minVersionIOS"`
	MinVersionAndroid   int   `json:"minVersionAndroid" bson:"minVersionAndroid"`
	Count               int64 `json:"count" bson:"count"`
	ForceVersionIOS     int   `json:"forceVersionIOS" bson:"forceVersionIOS"`
	ForceVersionAndroid int   `json:"forceVersionAndroid" bson:"forceVersionAndroid"`
}

type Version struct {
	MinVersionIOS       int `json:"minVersionIOS" bson:"minVersionIOS"`
	MinVersionAndroid   int `json:"minVersionAndroid" bson:"minVersionAndroid"`
	ForceVersionIOS     int `json:"forceVersionIOS" bson:"forceVersionIOS"`
	ForceVersionAndroid int `json:"forceVersionAndroid" bson:"forceVersionAndroid"`
}

type Count struct {
	Count int       `json:"count" bson:"count"`
	Time  time.Time `json:"time" bson:"time"`
}

type NotificationStart struct {
	Count int64 `json:"count" bson:"count"`
}

type OfferData struct {
	StartTime    time.Time     `json:"startTime" bson:"startTime"`
	EndTime      time.Time     `json:"endTime" bson:"endTime"`
	IsShowDialog bool          `json:"isShowDialog" bson:"isShowDialog"`
	CurrentOffer *SpecialOffer `json:"currentOffer" bson:"currentOffer"`
}

type UsernameError struct {
	UserID      int    `json:"userID" bson:"userID"`
	Username    string `json:"username" bson:"username"`
	SuggestName string `json:"suggestName" bson:"suggestName"`
}

type ErrorPurchase struct {
	UserID      int    `json:"userID" bson:"userID"`
	Error       string `json:"error" bson:"error"`
	ProductType string `json:"productType" bson:"productType"`
	ProductID   string `json:"productID" bson:"productID"`
}

type BetSpinner struct {
	ID  int `json:"id" bson:"id"`
	Bet int `json:"bet" bson:"bet"`
}

type BetResponse struct {
	WinChips    int              `json:"winChips" bson:"winChips"`
	Inventory   *PersonInventory `json:"inventory" bson:"inventory"`
	Indexes     []int            `json:"indexes" bson:"indexes"`
	SpinBuyInfo SpinBuyInfo      `json:"spinBuyInfo" bson:"spinBuyInfo"`
}

type SpinBuyInfo struct {
	IsBought bool `json:"isBought" bson:"isBought"`
	KWin     int  `json:"kWin" bson:"kWin"`
	LeftSpin int  `json:"leftSpin" bson:"leftSpin"`
}

type SpinCoeff struct {
	ID    int     `json:"id"`
	Index int     `json:"index"`
	Count int     `json:"count"`
	Coef  float64 `json:"coef"`
}
