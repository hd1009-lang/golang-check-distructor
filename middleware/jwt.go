package middleware

import (
	"director/configs"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	Phone      string
	SupplierId uint64
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.Config("SECRET")), nil
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)
	if err != nil {
		return nil, err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	data, ok := claims["data"].(map[string]interface{})
	supplierId := int64(data["supplier_id"].(float64))
	phone := data["phone"].(string)
	if !ok {
		return nil, errors.New("Something went wrong")
	}
	return &TokenPayload{
		SupplierId: uint64(supplierId),
		Phone:      string(phone),
	}, nil
}
