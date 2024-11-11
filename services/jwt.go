package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtsecret = "hellobsdk"
)

var jwtKey = []byte(jwtsecret)

type Claims struct {
	UserID               uint `json: "user_id"`
	jwt.RegisteredClaims      //jwt internal function
}

// generate jWT token for given userID
func GenerateJWTToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(48 * time.Hour) //set JWT to expire in 48 hrs
	claims := &Claims{
		UserID: userID, //user ki userID to be given here
		RegisteredClaims: jwt.RegisteredClaims{ //jwt lib internal function defining what claims to use and at what time it expires
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("jwtkey", token.SignedString(jwtKey))
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
