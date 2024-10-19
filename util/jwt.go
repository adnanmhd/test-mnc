package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"test-mnc/response"
	"time"
)

var secretKey = []byte("secret-key")

type CustomClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.MapClaims
}

func CreateToken(phoneNumber string) (tokenLogin response.Login, e error) {

	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		CustomClaims{
			PhoneNumber: phoneNumber,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return tokenLogin, err
	}
	tokenLogin.AccessToken = tokenString
	tokenLogin.RefreshToken = strconv.FormatInt(exp, 10)

	return tokenLogin, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ExtractPhoneNumber(tokenString string) (string, error) {
	var claims CustomClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil // Use your actual signing key here
	})

	if err != nil {
		return "", err
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.PhoneNumber, nil
}
