package entity

import (
	"WakaTImeGo/basic/database"
	"time"
)
import _ "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

type User struct {
	ID                string    `json:"id" gorm:"primary_key"`
	ApiKey            string    `json:"api_key" gorm:"unique; default:NULL"`
	Email             string    `json:"email" gorm:"index:idx_user_email; size:255"`
	Location          string    `json:"location"`
	Password          string    `json:"-"`
	CreatedAt         time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date" example:"2006-01-02 15:04:05.000"`
	LastLoggedInAt    time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date" example:"2006-01-02 15:04:05.000"`
	ShareDataMaxDays  int       `json:"-" gorm:"default:0"`
	ShareEditors      bool      `json:"-" gorm:"default:false; type:bool"`
	ShareLanguages    bool      `json:"-" gorm:"default:false; type:bool"`
	ShareProjects     bool      `json:"-" gorm:"default:false; type:bool"`
	ShareOSs          bool      `json:"-" gorm:"default:false; type:bool; column:share_oss"`
	ShareMachines     bool      `json:"-" gorm:"default:false; type:bool"`
	ShareLabels       bool      `json:"-" gorm:"default:false; type:bool"`
	IsAdmin           bool      `json:"-" gorm:"default:false; type:bool"`
	HasData           bool      `json:"-" gorm:"default:false; type:bool"`
	WakatimeApiKey    string    `json:"-"`
	WakatimeApiUrl    string    `json:"-"`
	SyncApiKey        string    `json:"-"` // api key for sync data
	ReportsWeekly     bool      `json:"-" gorm:"default:false; type:bool"`
	PublicLeaderboard bool      `json:"-" gorm:"default:false; type:bool"`
}

// Add add a user to database
func (user *User) Add() error {
	db := database.GetDb()
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Del delete a user form database
func (user *User) Del() error {
	db := database.GetDb()
	err := db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Update update a user in database
func (user *User) Update() error {
	db := database.GetDb()
	err := db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID get a user by id
func GetUserByID(id string) (User, error) {
	var user User
	db := database.GetDb()
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
