package v1

import "time"

/**
 * @Author: Kim
 * @Description: wakatime today fetch data
 * @File:  wakatime
 * @Date: 12/12/2022 3:35 PM
 */

type Today struct {
	CachedAt time.Time `json:"cached_at"`
	Data     TodayData `json:"data"`
}

type GrandTotal struct {
	Digital      string `json:"digital"`
	Hours        int    `json:"hours"`
	Minutes      int    `json:"minutes"`
	Text         string `json:"text"`
	TotalSeconds int    `json:"total_seconds"`
}

type Range struct {
	Date     time.Time `json:"date"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}

type TodayData struct {
	Categories       []interface{} `json:"categories"`
	Dependencies     []interface{} `json:"dependencies"`
	Editors          []interface{} `json:"editors"`
	Languages        []interface{} `json:"languages"`
	Machines         []interface{} `json:"machines"`
	OperatingSystems []interface{} `json:"operating_systems"`
	Projects         []interface{} `json:"projects"`
	GrandTotal       GrandTotal    `json:"grand_total"`
	Range            Range         `json:"range"`
}
