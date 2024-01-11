package service

import (
	"context"
	"github.com/Kazbekuly/grpc-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserServiceImpl(collection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{collection, ctx}
}

func (u *UserServiceImpl) FindUserById(id string) (*model.DbResponse, error) {
	userid, _ := primitive.ObjectIDFromHex(id)
	var user *model.DbResponse

	query := bson.M{"_id": userid}
	err := u.collection.FindOne(u.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.DbResponse{}, err
		}
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) FindUserByEmail(email string) (*model.DbResponse, error) {
	var user *model.DbResponse

	query := bson.M{"email": strings.ToLower(email)}
	err := u.collection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.DbResponse{}, err
		}
		return nil, err
	}
	return user, nil
}
