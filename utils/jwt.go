package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func CreateToken(email string, userId int64) (string, error)  {
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// func VerifyToken(token string) (int64, error) {
// 	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 	_, ok :=	token.Method.(*jwt.SigningMethodHMAC)
	
// 	if !ok {
// 		return nil, errors.New("Error happend in jwt parse")
// 	}
// 	return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		return 0, errors.New("Error in parse token")
// 	}

// 	tokenIsValid := parsedToken.Valid

// 	if !tokenIsValid {
// 		return 0, errors.New("Invalid Token")
// 	}

// 	claims, ok := parsedToken.Claims.(jwt.MapClaims)

// 	if !ok {
// 		return 0, errors.New("Invalid Token")
// 	}

// 	// email := claims["email"].(string)
// 	userId := int64(claims["userId"].(float64))

// 	return userId, nil
// }

func VerifyToken(tokenString string) (int64, error) {
    parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        log.Printf("Error parsing token: %v", err) // Log for debugging
        return 0, errors.New("invalid token")
    }

    if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
        userId, ok := claims["userId"].(float64)
        if !ok {
            return 0, errors.New("invalid token claims")
        }
        return int64(userId), nil
    }

    return 0, errors.New("invalid token")
}
