package service

import "github.com/Kazbekuly/grpc-crud/model"

type PostService interface {
	CreatePost(post *model.CreatePostRequest) (*model.DBPost, error)
	UpdatePost(string, model.UpdatePost) (*model.DBPost, error)
	FindPostById(string) (*model.DBPost, error)
	FindPosts(page int, limit int) (*[]model.DBPost, error)
	DeletePost(string) error
}
