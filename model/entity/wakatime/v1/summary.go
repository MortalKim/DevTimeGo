package v1

import "time"

/**
 * @Author: Kim
 * @Description: wakatime summaryService fetch data
 * @File:  wakatime
 * @Date: 12/12/2022 3:35 PM
 */

type Summary struct {
	CachedAt time.Time   `json:"cached_at"`
	Data     SummaryData `json:"data"`
}

type SummaryData struct {
	Categories       []SummaryItem `json:"categories"`
	Dependencies     []SummaryItem `json:"dependencies"`
	Editors          []SummaryItem `json:"editors"`
	GrandTotal       GrandTotal    `json:"grand_total"`
	Languages        []SummaryItem `json:"languages"`
	Machines         []SummaryItem `json:"machines"`
	OperatingSystems []SummaryItem `json:"operating_systems"`
	Projects         []SummaryItem `json:"projects"`
	Range            Range         `json:"range"`
}

type SummaryItem struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type GrandTotal struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Range struct {
	Date     string    `json:"date"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}
