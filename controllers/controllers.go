package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"user-management-api/entity"
	"user-management-api/domain"
	"user-management-api/repository"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
)

var userRepository repository.UserRepository

// Home - GET("/")
func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
}

// GetUser - GET("/user/{UserID}") - Return User as JSON object for given ID
func GetUser(context *gin.Context) {
	user, err := userRepository.GetUser(context.Param("UserID"))
	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

//AddUser - POST("/user") - Add user as per POST body
func AddUser(context *gin.Context) {
	var user domain.RequestBody
	context.BindJSON(&user)

	createdUser, err := createUser(user)

	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	err = userRepository.AddUser(*createdUser)

	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": createdUser})
}

//AuthenticateUser - Checks if given User ID and password are valid
func AuthenticateUser(context *gin.Context) {

	user, err := userRepository.GetUser(context.Param("UserID"))
	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(context.Param("Password")))
	if(err != nil) {
		context.JSON(http.StatusUnauthorized, gin.H{"data": "Authentication failed"})
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func createUser(user domain.RequestBody) (*entity.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if(err != nil) {
		return nil, err
	}

	return &entity.User{
		UserID: uuid.UUID.String(uuid.NewV4()),
		Name: user.Name,
		Password: string(password),
	}, nil
}

func init() {
	userRepository = repository.GetRepository()
}
