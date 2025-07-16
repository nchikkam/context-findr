package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nchikkam/context-findr-be/model"
	"github.com/nchikkam/context-findr-be/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// @Summary new user registration to the system
// @Description allows new user to register in system
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegistrationInput true "New user data"
// @Success 201 {object} model.RegistrationOrSigninResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input model.RegistrationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data for New User Registration"})
		return
	}

	var existingUser model.User
	err := getUserCollection().FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duplicate user creation"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := model.User{
		ID:       primitive.NewObjectID(),
		Name:     input.Name,
		Email:    strings.ToLower(input.Email),
		Password: string(hashedPassword),
	}

	_, err = getUserCollection().InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	token, err := utils.GenerateJWT(newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, model.RegistrationOrSigninResponse{
		Token: token,
		User: model.UserResponse{
			ID:    newUser.ID,
			Name:  newUser.Name,
			Email: newUser.Email,
		},
	})
}

// @Summary Signin a user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body model.SigninInput true "Signing in credentials"
// @Success 200 {object} model.RegistrationOrSigninResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signin [post]
func Signin(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input model.SigninInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var foundUser model.User
	err := getUserCollection().FindOne(ctx, bson.M{"email": input.Email}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials. Please check your email and password"})
		return
	}

	token, err := utils.GenerateJWT(foundUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, model.RegistrationOrSigninResponse{
		Token: token,
		User: model.UserResponse{
			ID:    foundUser.ID,
			Name:  foundUser.Name,
			Email: foundUser.Email,
		},
	})
}
