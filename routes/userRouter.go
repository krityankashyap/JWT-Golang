package routes

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/krityan/golang-jwt/middlewares"
	controller "github.com/krityan/golang-jwt/controllers"
)

func UserRoute(incomingRoute *gin.Engine) {
	incomingRoute.Use(middleware.Authenticate);

	incomingRoute.GET("/user", controller.GetUsers());

	incomingRoute.GET("/user/:user_id", controller.GetUser());
}