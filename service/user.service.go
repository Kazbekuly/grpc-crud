package service

import "github.com/Kazbekuly/grpc-crud/model"

type UserService interface {
	FindUserById(id string) (*model.DbResponse, error)
	FindUserByEmail(email string) (*model.DbResponse, error)
}
