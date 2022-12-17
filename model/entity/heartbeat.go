package entity

import (
	"WakaTImeGo/basic/database"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type Heartbeat struct {
	ID              uint64    `json:"-" gorm:"primary_key; auto_increment"`
	User            *User     `json:"-" gorm:"not null; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" hash:"ignore"`
	UserID          string    `json:"-" gorm:"not null; index:idx_time_user; index:idx_user_project"` // idx_user_project is for quickly fetching a userService's project list (settings page)
	Entity          string    `json:"entity" gorm:"not null"`
	Type            string    `json:"type" gorm:"size:255"`
	Category        string    `json:"category" gorm:"size:255"`
	Project         string    `json:"project" gorm:"index:idx_project; index:idx_user_project"`
	Branch          string    `json:"branch" gorm:"index:idx_branch"`
	Language        string    `json:"language" gorm:"index:idx_language"`
	IsWrite         bool      `json:"is_write"`
	Editor          string    `json:"editor" gorm:"index:idx_editor" hash:"ignore"`                     // ignored because editor might be parsed differently by wakatime
	OperatingSystem string    `json:"operating_system" gorm:"index:idx_operating_system" hash:"ignore"` // ignored because os might be parsed differently by wakatime
	Machine         string    `json:"machine" gorm:"index:idx_machine" hash:"ignore"`                   // ignored because wakatime api doesn't return machines currently
	UserAgent       string    `json:"user_agent" hash:"ignore" gorm:"type:varchar(255)"`
	Time            time.Time `json:"time" gorm:"type:timestamp(3); index:idx_time; index:idx_time_user" swaggertype:"primitive,number"`
	Hash            string    `json:"-" gorm:"type:varchar(17); uniqueIndex"`
	IsCounted       bool      `json:"-" gorm:"column:is_counted; default:false"` // whether this heartbeat has been counted in the user's duration
	Origin          string    `json:"-" hash:"ignore" gorm:"type:varchar(255)"`
	OriginId        string    `json:"-" hash:"ignore" gorm:"type:varchar(255)"`
	CreatedAt       time.Time `json:"created_at" gorm:"type:timestamp(3)" hash:"ignore"` // https://gorm.io/docs/conventions.html#CreatedAt
}

// Add heartbeat to database
func (heartbeat *Heartbeat) Add() error {
	db := database.GetDb()
	err := db.Create(heartbeat).Error
	if err != nil {
		return err
	}
	return nil
}

func (heartbeat *Heartbeat) HashSelf() {
	md := md5.New()
	md.Write([]byte(heartbeat.UserID + heartbeat.Time.String() + heartbeat.Entity + heartbeat.Type + heartbeat.Category + heartbeat.Project + heartbeat.Branch + heartbeat.Language + heartbeat.Editor + heartbeat.OperatingSystem + heartbeat.Machine + heartbeat.UserAgent))
	heartbeat.Hash = hex.EncodeToString(md.Sum(nil))[8:24]
}

// Update update heartbeat
func (heartbeat *Heartbeat) Update() error {
	db := database.GetDb()
	err := db.Save(heartbeat).Error
	if err != nil {
		return err
	}
	return nil
}
