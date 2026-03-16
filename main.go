package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/krityan/golang-jwt/routes"
)


func main() {
	port:= os.Getenv("PORT");

	if port == "" {
		port = "8000"
	}

	router:= gin.New();  // New returns a new blank Engine instance without any middleware attached.

	router.Use(gin.Logger());  // Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.

	routes.AuthRoute(router);

	routes.UserRoute(router);

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	});

	router.GET("api-2", func(c *gin.Context) {
		c.JSON(200 , gin.H{"sucess": "Access granted for api-2"});
	});

	router.Run(":" + port);


}