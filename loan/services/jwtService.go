package services

import (
	"context"
	"loan/models"
	"loan/proto"
	"log"
)

func ValidateJwt(authRqst models.AuthRequest) bool {
	grpcConn := GrpcConn()

	a := proto.NewAuthServiceClient(grpcConn)

	authResponse, authRespErr := a.ValidateJWT(context.Background(), &proto.AuthRequest{
		JwtToken: authRqst.JWTToken,
	})

	if authRespErr != nil {
		log.Println("authRespErr", authRespErr)
	}

	log.Println("authResponse", authResponse)
	return authResponse.Valid
}
