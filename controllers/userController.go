package controllers

import (
 "context"
 "fmt"
 "log"
 "net/http"
 "time"
 "github.com/gin-gonic/gin"
 "github.com/go-playground/validator/v10"
 utils "github.com/krityan/golang-jwt/utils"
 "github.com/krityan/golang-jwt/models"
  database "github.com/krityan/golang-jwt/config"
 "golang.org/x/crypto/bcrypt"

)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user");
var Validate= validator.New();  // New returns a new instance of 'validate' with sane defaults.
 
func HashPassword() 

func VerifyPassword()

func SignUp()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc {  // Only admin can check the details of other users
 return func(c *gin.Context) {
	UserId := c.Param("user_id");

	err := utils.MatchUserTypeToUId(UserId, c);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
 }
 
}