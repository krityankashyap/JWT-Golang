package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(role string, c *gin.Context) (err error) {

	userType:= c.GetString("user_type");
	err = nil;

	if userType!= role {
		err= errors.New("Unauthorized to access this resource");
		return err;
	}
	return err;
}

func MatchUserTypeToUId(userId string, c *gin.Context) (err error) {
  
	userType := c.GetString("user_type");
	uId := c.GetString("uid");
  err= nil;

	if userType== "USER" && uId != userId {
		err = errors.New("Unauthorized to access this resource");
		return err;
	}

	err = CheckUserType(userType, c);
	return err;
}