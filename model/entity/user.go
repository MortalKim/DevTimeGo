package entity

import (
	"WakaTImeGo/basic/database"
	"gorm.io/gorm"
	"time"
)

import _ "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

type User struct {
	ID        uint           `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp;default:NULL"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:timestamp;default:NULL"`

	ApiKey            string    `json:"api_key" gorm:"unique; default:NULL"`
	UserName          string    `json:"user_name" gorm:"default:NULL"`
	Email             string    `json:"email" gorm:"index:idx_user_email; size:255"`
	Birthday          time.Time `json:"birthday" gorm:"default:NULL"`
	Gender            string    `json:"gender" gorm:"default:NULL"`
	Location          string    `json:"location"`
	Password          string    `json:"-"`
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
	ReportsWeekly     bool      `json:"-" gorm:"default:false; type:bool"`
	PublicLeaderboard bool      `json:"-" gorm:"default:false; type:bool"`
}

// Add add a userController to database
func (user *User) Add() error {
	db := database.GetDb()
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Del delete a userController form database
func (user *User) Del() error {
	db := database.GetDb()
	err := db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Update update a userController in database
func (user *User) Update() error {
	db := database.GetDb()
	err := db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID get a userController by id
func GetUserByID(id string) (User, error) {
	var user User
	db := database.GetDb()
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByApiKey(apiKey string) (User, error) {
	var user User
	db := database.GetDb()
	err := db.Where("api_key = ?", apiKey).First(&user).Error
	return user, err
}
