package controllers

import (
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

	createdUser := createUser(user)

	err := userRepository.AddUser(createdUser)

	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": createdUser})
}

func createUser(user domain.RequestBody) entity.User{
	return entity.User{UserID: uuid.UUID.String(uuid.NewV4()),Name: user.Name}
}

func init() {
	userRepository = repository.GetRepository()
}
