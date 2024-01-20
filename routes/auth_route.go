package routes

import (
	"github.com/Kazbekuly/grpc-crud/handler"
	"github.com/Kazbekuly/grpc-crud/service"
	"github.com/gin-gonic/gin"
)

type AuthControllerRoute struct {
	authController handler.AuthController
}

func NewAuthControllerRoute(authControllerRoute handler.AuthController) AuthControllerRoute {
	return AuthControllerRoute{authControllerRoute}
}

func (ar *AuthControllerRoute) AuthRoute(rg *gin.RouterGroup, userService service.AuthService) {
	router := rg.Group("/auth")
	router.POST("/register", ar.authController.SignUpUser)
	router.POST("/login", ar.authController.SignInUser)
	router.GET("/refresh")
	router.GET("/logout")
}
