package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SignUpInput struct {
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	PasswordConfirm    string    `json:"passwordConfirm"`
	Role               string    `json:"role"`
	VerificationCode   string    `json:"verificationCode"`
	ResetPasswordToken string    `json:"resetPasswordToken"`
	ResetPasswordAt    string    `json:"resetPasswordAt"`
	Verified           bool      `json:"verified"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DbResponse struct {
	Id                 primitive.ObjectID `json:"id" bson:"id"`
	Name               string             `json:"name" bson:"name"`
	Email              string             `json:"email" bson:"email"`
	Password           string             `json:"password" bson:"password"`
	PasswordConfirm    string             `json:"passwordConfirm" bson:"passwordConfirm,omitempty"`
	Role               string             `json:"role" bson:"role"`
	VerificationCode   string             `json:"verificationCode,omitempty" bson:"verificationCode"`
	ResetPasswordToken string             `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ResetPasswordAt    string             `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	Verified           bool               `json:"verified" bson:"verified"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type UserResponse struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func FilteredResponse(user *DbResponse) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
