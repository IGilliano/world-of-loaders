package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"worldOfLoaders/pkg/models"
	"worldOfLoaders/pkg/repository"
)

const (
	salt       = "1gdfg734tybs"
	signingKey = "df2154gs365661sd"
	tokenTTL   = 200 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Id    int    `json:"id"`
	Class string `json:"class"`
}

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (a *AuthService) CreatePlayer(player models.Player) (int, error) {
	player.Password = generatePasswordHash(player.Password)
	return a.rep.CreatePlayer(player)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthService) GetPlayers() ([]*models.Player, error) {
	return a.rep.GetPlayers()
}

func (a *AuthService) GenerateToken(login, password string) (string, error) {
	player, err := a.rep.GetPlayer(login, generatePasswordHash(password))
	if err != nil {
		fmt.Println("Error! Incorrect login or password")
		return "", err
	}

	claims := &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		},
		Id:    player.ID,
		Class: player.Class,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, claims.Class, nil
}
