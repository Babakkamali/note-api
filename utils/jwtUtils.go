package utils

import (
    "time"
    "os"

    "github.com/dgrijalva/jwt-go"
	
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if !token.Valid {
        return nil, jwt.NewValidationError("token has expired or is not valid", jwt.ValidationErrorExpired)
    }

    return claims, err
}
