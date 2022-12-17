package request

/**
 * @Author: Kim
 * @Description:
 * @File:  user
 * @Date: 12/15/2022 7:32 PM
 */

// UserRegister user register prams struct
type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
