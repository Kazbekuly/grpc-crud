package service

import (
	"context"
	"errors"
	"github.com/Kazbekuly/grpc-crud/model"
	"github.com/Kazbekuly/grpc-crud/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthServiceImpl(collection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{collection, ctx}
}

func (a *AuthServiceImpl) SignUpUser(input *model.SignUpInput) (*model.DbResponse, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = input.CreatedAt
	input.Email = strings.ToLower(input.Email)
	input.PasswordConfirm = ""
	input.Verified = true
	input.Role = "user"

	hashedPassword, _ := utils.HashPassword(input.Password)
	input.Password = hashedPassword
	res, err := a.collection.InsertOne(a.ctx, &input)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exists")
		}
		return nil, err
	}
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := a.collection.Indexes().CreateOne(a.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}
	var newUser *model.DbResponse
	query := bson.M{"_id": res.InsertedID}
	err = a.collection.FindOne(a.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, err
}

func (a *AuthServiceImpl) SignInUser(input *model.SignInInput) (*model.DbResponse, error) {
	return nil, nil
}
