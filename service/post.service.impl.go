package service

import (
	"context"
	"errors"
	"github.com/Kazbekuly/grpc-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PostServiceImpl struct {
	postCollection *mongo.Collection
	ctx            context.Context
}

func NewPostService(postCollection *mongo.Collection, ctx context.Context) *PostServiceImpl {
	return &PostServiceImpl{postCollection, ctx}
}

func (p *PostServiceImpl) CreatePost(post *model.CreatePostRequest) (*model.DBPost, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt
	res, err := p.postCollection.InsertOne(p.ctx, post)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("post already exists")
		}
		return nil, err
	}
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := p.postCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for title")
	}
	var newPost *model.DBPost
	query := bson.M{"_id": res.InsertedID}
	if err = p.postCollection.FindOne(p.ctx, query).Decode(&newPost); err != nil {
		return nil, err
	}
	return newPost, nil
}
