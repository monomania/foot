package vo

type BattleData struct {
	Year         string `json:"Year"`
	Date         string `json:"Date"`
	League       string `json:"League"`
	Home         string `json:"Home"`
	Guest        string `json:"Guest"`
	HomeTeamID   string `json:"HomeTeamId"`
	GuestTeamID  string `json:"GuestTeamId"`
	FT           string `json:"FT"`
	HT           string `json:"HT"`
	HColNo       int    `json:"HColNo"`
	GColNo       int    `json:"GColNo"`
	HhalfColNo   int    `json:"HhalfColNo"`
	GhalfColNo   int    `json:"GhalfColNo"`
	Result       string `json:"Result"`
	Odds         string `json:"Odds"`
	SclassID     string `json:"SclassId"`
	ResultHalf   string `json:"ResultHalf"`
	LetgoalHalf  string `json:"LetgoalHalf"`
	OuResult     string `json:"OuResult"`
	Ou           string `json:"Ou"`
	OuResultHalf string `json:"OuResultHalf"`
	OuHalf       string `json:"OuHalf"`
}
