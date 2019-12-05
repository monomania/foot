package vo

type AsiaData struct {
	ScheduleID int `json:"scheduleId"`
	Companies  []struct {
		CompanyID int    `json:"companyId"`
		NameCn    string `json:"nameCn"`
		NameEn    string `json:"nameEn"`
		Details   []struct {
			OddsID        int     `json:"oddsId"`
			Num           int     `json:"num"`
			FirstHomeOdds float64 `json:"firstHomeOdds"`
			FirstDrawOdds float64 `json:"firstDrawOdds"`
			FirstAwayOdds float64 `json:"firstAwayOdds"`
			HomeOdds      float64 `json:"homeOdds"`
			DrawOdds      float64 `json:"drawOdds"`
			AwayOdds      float64 `json:"awayOdds"`
			ModifyTime    string  `json:"modifyTime"`
		} `json:"details"`
	} `json:"companies"`
}
