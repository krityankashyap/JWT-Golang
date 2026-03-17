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
 "go.mongodb.org/mongo-driver/bson"

)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user");
var Validate= validator.New();  // New returns a new instance of 'validate' with sane defaults.
 
func HashPassword() 

func VerifyPassword()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel= context.WithTimeout(context.Background(), 100*time.Second);

		var user models.User;  // SignUp functions creates the USER in the database and returns the user details in the response

		err := c.BindJSON(&user); // BindJSON binds the received JSON to the user variable
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
			return;
		}

		validateError := Validate.Struct(user); // Validate.Struct validates the struct fields based on the tags defined in the struct

		if validateError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateError.Error()});
			return;
		}

		// count of the user if count of provided email or phone number is greater than 0 then the email or phone number is already in use

		count , err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email}); // CountDocuments returns the count of the documents that match the filter
		defer cancel();
		if err != nil {
			log.Panic(err);
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking the email"});
			return;
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone}); // CountDocuments returns the count of the documents that match the filter

		defer cancel();
		if err != nil {
			log.Panic(err);
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking the phone number"});
			return;
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H("error": "Email or phone number already in use"));
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)); // Parse parses a formatted string and returns the time value it represents

		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

		user.ID = primitive.NewObjectID(); // NewObjectID returns a new ObjectID generated from the current timestamp, machine id, process id, and a random counter.

		user.User_id= user.ID.Hex(); // Hex returns the hex string reprsentation of the object Id

		token, refreshToken, _ := utils.GenerateToken(*user.Email, *user.First_name, *user.Last_name, *&user.User_type, *&user.User_id); // GenerateToken generates the JWT token and refresh token for the user

		user.Token = &token;
		user.Refresh_token = &refreshToken;

		// We have to insert the user in the database after hashing the password and generating the token and refresh token

		resultInsertionNumber, insertErr := UserCollection.InsertOne(ctx, user); // InsertOne inserts a single document into the collection and returns the result of the insertion

		if insertErr != nil {
			log.Panic(insertErr);
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while inserting the user"});
			return;
		}

		defer cancel();
		c.JSON(http.StatusOK, resultInsertionNumber); // if the user is inserted successfully then return the result of the insertion
	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc {  
 return func(c *gin.Context) {
	UserId := c.Param("user_id");

	err := utils.MatchUserTypeToUId(UserId, c);  // Only admin can check the details of other users
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second);

	var user models.User;
	err = UserCollection.FindOne(ctx, bson.M{"user_id": UserId}).Decode(&user); // FindOne returns a SingleResult which is decoded into the user variable

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the user details"});
		return;
	}

	defer cancel();
	c.JSON(http.StatusOK, user); // if the user is found then return the user details
 }
 
}