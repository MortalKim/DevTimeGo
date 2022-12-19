package entity

import "github.com/golang-jwt/jwt/v4"

/**
 * @Author: Kim
 * @Description:
 * @File:  Jwt
 * @Date: 12/19/2022 5:40 PM
 */

// JwtStruct struct used for jwt
type JwtStruct struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
