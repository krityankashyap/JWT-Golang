package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/krityan/golang-jwt/controllers"
)

func AuthRoute(incomingRoute *gin.Engine) {
	incomingRoute.POST("/user/signup", controller.SignUp);
	incomingRoute.POST("/user/login", controller.Login);
}