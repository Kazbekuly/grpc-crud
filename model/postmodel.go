package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreatePostRequest struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Text      string    `json:"text" bson:"text" binding:"required"`
	Image     string    `json:"image" bson:"image" binding:"required"`
	User      string    `json:"user" bson:"user" binding:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}

type DBPost struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Text      string             `json:"text,omitempty" bson:"text,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	User      string             `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}

type UpdatePost struct {
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Text      string    `json:"text,omitempty" bson:"text,omitempty"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"`
	User      string    `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}
