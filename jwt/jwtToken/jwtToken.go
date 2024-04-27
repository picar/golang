package jwtToken

import (
	"github.com/golang-jwt/jwt/v5"
 "time"
)

func CreateToken(username string, secret string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
    return "", err
    }

 return tokenString, nil
}