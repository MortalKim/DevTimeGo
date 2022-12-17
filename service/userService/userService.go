package userService

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
	//check if user admin is exist
	_, err := entity.GetUserByID("1")
	//if not exist, create it
	if err != nil {
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
}

// CanRegister Can user register
func CanRegister(username, email string) (bool, string) {
	_, err := GetUserByUserName(username)
	if err == nil {
		return false, "username is exist"
	}
	_, err = GetUserByEmail(email)
	if err == nil {
		return false, "email is exist"
	}
	return true, ""
}

// Register Register a new user
func Register(username, email, password string) (entity.User, error) {
	var user entity.User
	user.UserName = username
	user.Email = email
	user.Password, _ = bcrypt.PwdHash(password)
	user.ApiKey = uuid.New().String()
	err := user.Add()
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUserByApiKey Get userService by ApiKey
func GetUserByApiKey(apiKey string) (entity.User, error) {
	return entity.GetUserByApiKey(apiKey)
}

// GetUserById Get user by ID
func GetUserById(id string) (entity.User, error) {
	return entity.GetUserByID(id)
}

// GetUserByEmail Get user by email
func GetUserByEmail(email string) (entity.User, error) {
	return entity.GetUserByEmail(email)
}

// GetUserByUserName Get user by username
func GetUserByUserName(username string) (entity.User, error) {
	return entity.GetUserByUserName(username)
}
