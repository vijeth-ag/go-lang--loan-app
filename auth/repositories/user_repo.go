package repositories

import (
	"auth/model"
	"auth/utils"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (userRepo *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	log.Println("at CreateUser", user)
	var passwordHashErr error
	user.Password, passwordHashErr = utils.HashPassword(user.Password)
	if passwordHashErr != nil {
		log.Println("err hashsing password")
	}

	res, err := userRepo.db.Collection("users").InsertOne(ctx, user)
	log.Println("err", err)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			log.Println("User exists")
		}
		return errors.New("user exists")
	}
	log.Println("insert ok", res)
	return nil
}

func (UserRepo *UserRepo) UpdateUser(userName string, updateStr string) error {

	var userUpdate interface{}

	json.Unmarshal([]byte(updateStr), &userUpdate)

	update := bson.M{
		"$set": userUpdate,
	}

	log.Println("the update", update)

	filter := bson.D{{Key: "userName", Value: userName}}

	log.Println("filter", filter)

	res, err := UserRepo.db.Collection("users").UpdateOne(context.TODO(), filter, update)
	log.Println("res", res)

	if err != nil {
		log.Println("err", err)
		return err
	}
	return nil
}
