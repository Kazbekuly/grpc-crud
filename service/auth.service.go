package service

import "github.com/Kazbekuly/grpc-crud/model"

type AuthService interface {
	SignUpUser(input *model.SignUpInput) (*model.DbResponse, error)
	SignInUser(input *model.SignInInput) (*model.DbResponse, error)
}
