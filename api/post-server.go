package api

import (
	pb "github.com/Kazbekuly/grpc-crud/pb"
	"github.com/Kazbekuly/grpc-crud/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostServer struct {
	pb.UnimplementedPostServiceServer
	postCollection *mongo.Collection
	postService    service.PostService
	authService    service.AuthService
	userService    service.UserService
	userCollection *mongo.Collection
}

func NewGrpcPostServer(postCollection *mongo.Collection, postService service.PostService) (*PostServer, error) {
	postServer := &PostServer{
		postCollection: postCollection,
		postService:    postService,
	}
	return postServer, nil
}
