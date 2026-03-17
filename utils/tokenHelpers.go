package utils

import (

	database "github.com/krityan/golang-jwt/config"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type SignedDetails struct {
	Email            string
	First_name       string
	Last_name        string
	Uid              string
	User_type        string
	jwt.StandardClaims
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user");

var SECRET_KEY= os.Getenv("JWT_SECRET");

func GenerateToken(email string, firstName string, lastName string, uid string, userType string) (signedToken string, refreshedSignedToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		First_name: firstName,
		Last_name: lastName,
		Uid: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token , err := jwt.NewWithClaims(jwt.SigningMethodhs256 , claims).SignedString([]byte(SECRET_KEY));
	refreshToken , err := jwt.NewWithClaims(jwt.SigningMethodhs256, refreshClaims).SignedString([]byte(SECRET_KEY));

	if err != nil {
		log.Panic(err);
		return;
	}

	return token, refreshToken, err;
}