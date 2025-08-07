package model

type StatisticsNew struct {
	WinLose     *ComparisonModel `json:"winLose" bson:"winLose"`
	TillendPass *ComparisonModel `json:"tillendPass" bson:"tillendPass"`
	ThrowTaken  *ComparisonModel `json:"throwTaken" bson:"throwTaken"`
}

type ComparisonModel struct {
	Positive int `json:"positive" bson:"positive"`
	Negative int `json:"negative" bson:"negative"`
}
