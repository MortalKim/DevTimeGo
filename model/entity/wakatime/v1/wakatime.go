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

type TodayData struct {
	Categories       []Categories       `json:"categories"`
	Dependencies     []Dependencies     `json:"dependencies"`
	Editors          []Editors          `json:"editors"`
	GrandTotal       GrandTotal         `json:"grand_total"`
	Languages        []Languages        `json:"languages"`
	Machines         []Machines         `json:"machines"`
	OperatingSystems []OperatingSystems `json:"operating_systems"`
	Projects         []Projects         `json:"projects"`
	Range            Range              `json:"range"`
}

type Categories struct {
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
type Dependencies struct {
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
type Editors struct {
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
type Languages struct {
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
type Machines struct {
	Decimal       string  `json:"decimal"`
	Digital       string  `json:"digital"`
	Hours         int     `json:"hours"`
	MachineNameID string  `json:"machine_name_id"`
	Minutes       int     `json:"minutes"`
	Name          string  `json:"name"`
	Percent       float64 `json:"percent"`
	Seconds       int     `json:"seconds"`
	Text          string  `json:"text"`
	TotalSeconds  float64 `json:"total_seconds"`
}
type OperatingSystems struct {
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
type Projects struct {
	Color        interface{} `json:"color"`
	Decimal      string      `json:"decimal"`
	Digital      string      `json:"digital"`
	Hours        int         `json:"hours"`
	Minutes      int         `json:"minutes"`
	Name         string      `json:"name"`
	Percent      float64     `json:"percent"`
	Seconds      int         `json:"seconds"`
	Text         string      `json:"text"`
	TotalSeconds float64     `json:"total_seconds"`
}
type Range struct {
	Date     string    `json:"date"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}
