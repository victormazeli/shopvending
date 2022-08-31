package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserId    int    `json:"user_id"`
	UserType  string `json:"user_type"`
	IsRefresh bool   `json:"is_refresh"`
	jwt.StandardClaims
}

type JwtService struct{}

func (j JwtService) GenerateTokenPair(userId int, userType string) (map[string]string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &JWTClaim{
		UserId:    userId,
		UserType:  userType,
		IsRefresh: false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	refreshExpirationTime := time.Now().Add(48 * time.Hour)
	refreshClaims := &JWTClaim{
		UserId:    userId,
		IsRefresh: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	r, er := refreshToken.SignedString(jwtKey)
	if er != nil {
		return nil, er
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": r,
	}, nil
}

func (j JwtService) ValidateToken(signedToken string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("invalid token")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	payload := make(map[string]interface{})

	payload["sub"] = struct {
		UserId    int  `json:"user_id"`
		IsRefresh bool `json:"user_type"`
	}{claims.UserId, claims.IsRefresh}
	return claims.UserId, err
}
