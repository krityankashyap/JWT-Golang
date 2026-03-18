package utils

import (
	"context"
	"log"
	"os"
	"time"

	database "github.com/krityan/golang-jwt/config"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	token , err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY));
	refreshToken , err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY));

	if err != nil {
		log.Panic(err);
		return;
	}

	return token, refreshToken, err;
}

func UpdateAllTokens(signedToken string, refreshToken string, user_id string){
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second);

	var updateObj primitive.D; // primitive.D is a BSON document that represents a single key-value pair. It is used to specify the update operations to be performed on the document.

	updateObj = append(updateObj, bson.E{"token", signedToken});
	updateObj = append(updateObj, bson.E{"refresh_token", refreshToken});

	Updated_at,_ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

	updateObj = append(updateObj, bson.E{"updated_at", Updated_at});

	upsert := true; // upsert is a boolean value that specifies whether to insert a new document if no document matches the filter. If upsert is true, a new document will be inserted if no document matches the filter. If upsert is false, no document will be inserted if no document matches the filter.

	filter := bson.M{"user_id": user_id}; // bson.M is a BSON document that represents a map of key-value pairs. It is used to specify the filter criteria for the update operation.
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			bson.E{"$set", updateObj},
		},
		&opt,
	)

	defer cancel();

	if err != nil {
		log.Panic(err);
		return;
	}
	return;
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token , err := jwt.ParseWithClaims(signedToken, 
		                                &SignedDetails{}, 
																		func(token *jwt.Token) (interface{}, error) {
																			return []byte(SECRET_KEY), nil;
																		},

	)

	if err != nil {
		msg = err.Error();
		return;
	}

	claims, ok := token.Claims.(*SignedDetails);
	if !ok {
		msg = "the token is invalid"
		msg = err.Error();
		return;
	}

	return claims, msg;
}