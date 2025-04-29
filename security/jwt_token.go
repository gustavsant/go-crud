package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserEmail string `json:"user_email"`
	jwt.StandardClaims
}

var jwtKey = RetrieveSecurityToken()

func GenerateAndSignJWT(userEmail string) (string, error) {
	expDate := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expDate.Unix(),
			Issuer:    "go-crud",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	fmt.Println([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected token signining method")
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
