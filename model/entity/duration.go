package entity

import (
	"WakaTImeGo/basic/database"
	"time"
)

/**
 * @Author: Kim
 * @Description:
	form wakatime:
		A user's coding activity for the given day as an array of durations.
		Durations are read-only representations of Heartbeats,
		created by joining multiple Heartbeats together when they’re within 15 minutes of each other.
		The 15 minutes default can be changed with your account’s Keystroke Timeout preference.
	by this project:
		duration timeout is 5 minutes by default, and now there is no settings to change it.
 * @File:  duration
 * @Date: 12/13/2022 5:12 PM
*/

type Duration struct {
	ID              uint64        `json:"-" gorm:"primary_key; auto_increment"`
	UserID          string        `json:"user_id" gorm:"not null; index:idx_time_user; index:idx_user_project"`
	Time            time.Time     `json:"time" hash:"ignore"`
	Duration        time.Duration `json:"duration" hash:"ignore"`
	Project         string        `json:"project"`
	Language        string        `json:"language"`
	Editor          string        `json:"editor"`
	OperatingSystem string        `json:"operating_system"`
	Machine         string        `json:"machine"`
	Branch          string        `json:"branch"`
	NumHeartbeats   int           `json:"-" hash:"ignore"`
}

//Add a new Duration to database
func (d *Duration) Add() error {
	db := database.GetDb()
	err := db.Create(d).Error
	if err != nil {
		return err
	}
	return nil
}

//Update a Duration in database
func (d *Duration) Update() error {
	db := database.GetDb()
	err := db.Save(d).Error
	if err != nil {
		return err
	}
	return nil
}
