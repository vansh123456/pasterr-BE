package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("heyheybsdk")

type Claims struct {
	UserID uint `json:"user_id"` // Corrected json tag
	jwt.RegisteredClaims
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
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
	fmt.Println("Parsing JWT:", tokenStr) // Add more debug prints
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err) // Print specific error
		return nil, err
	}

	if !token.Valid {
		fmt.Println("Token is not valid")
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
