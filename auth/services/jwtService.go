package services

import (
	"auth/model"
	"auth/proto"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type MyGrpcServer model.GrpcServer

var jwtSecret = []byte("")

func init() {
	log.Println("loading env")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("err loading env")
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	log.Println("--", string(jwtSecret))
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiredAt"`
}

func NewPayload(username string) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		log.Println("err generating tokenId")
	}
	log.Println("tokenId", tokenId)

	payload := &Payload{
		ID:        tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 100),
	}

	return payload, nil
}

var ErrExpiredToken = errors.New("token has expired")

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}

func GenerateJWTToken(username string) (string, time.Time) {
	log.Println("username", username)

	payload, _ := NewPayload(username)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	log.Println("jwtSecret", jwtSecret)
	tokenStr, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("err", err)
	}
	log.Println("tokenStr", tokenStr)
	return tokenStr, payload.ExpiresAt
}

func (g MyGrpcServer) ValidateJWT(ctx context.Context, authRqst *proto.AuthRequest) (*proto.AuthResponse, error) {
	log.Println("at jwtService someone called ValidateJWT via grpc", authRqst)

	payload := &Payload{}

	log.Println("jwtSecret:::::::::", jwtSecret)

	tkn, err := jwt.ParseWithClaims(authRqst.JwtToken, payload, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	log.Println("errx", err)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("ErrSignatureInvalid")
			return &proto.AuthResponse{
				Valid: false,
			}, nil
		}
	}

	if !tkn.Valid {
		log.Println("tkn.notValid ", tkn.Valid)
		return &proto.AuthResponse{
			Valid: false,
		}, nil
	}

	return &proto.AuthResponse{
		Valid: true,
	}, nil

}
