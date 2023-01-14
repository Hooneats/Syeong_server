package ciper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"time"
)

type Token struct {
	AccessToken  TokenDetail
	RefreshToken TokenDetail
}

type TokenDetail struct {
	UUID     string
	Token    string
	Duration int64
}

func CreateToken(userID, accessKey, refreshKey string) (*Token, error) {
	//Creating Access Token
	acToken, err := createAccessToken(userID, accessKey)
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rfToken, err := createRefreshToken(userID, refreshKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  *acToken,
		RefreshToken: *rfToken,
	}, nil
}

func createAccessToken(userID, accessKey string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.Duration = time.Now().Add(time.Minute * 15).Unix()
	ati, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	td.UUID = ati.String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.UUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.Duration
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.Token, err = at.SignedString([]byte(accessKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func createRefreshToken(userID, refreshKey string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.Duration = time.Now().Add(time.Hour * 2).Unix()
	rti, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	td.UUID = rti.String()

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.UUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.Duration
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.Token, err = rt.SignedString([]byte(refreshKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}
