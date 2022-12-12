package user

import (
	"WakaTImeGo/model/entity"
	"WakaTImeGo/utils/bcrypt"
)
import "WakaTImeGo/basic/database"
import "github.com/google/uuid"

func InitDatabase() {
	//create table to database
	err := database.GetDb().AutoMigrate(&entity.User{})
	if err != nil {
		return
	}
	CreateAdminUser()
}

// CreateAdminUser create userController named admin
func CreateAdminUser() {
	var user entity.User
	user.UserName = "admin"
	user.Email = "admin@admin"
	user.Password, _ = bcrypt.PwdHash("admin")
	user.IsAdmin = true
	user.ApiKey = uuid.New().String()
	err := user.Add()
	if err != nil {
		return
	}
}
