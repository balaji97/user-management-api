package controllers

import (
	"user-management-api/entity"
	"user-management-api/domain"
	"user-management-api/repository"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
)

// Home - GET("/")
func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
}

// GetUser - GET("/user/{UserID}") - Return User as JSON object for given ID
func GetUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": repository.GetUser(context.Param("UserID"))})
}

//AddUser - POST("/user") - Add user as per POST body
func AddUser(context *gin.Context) {
	var user domain.RequestBody
	context.BindJSON(&user)

	createdUser := createUser(user)

	repository.AddUser(createdUser)
	context.JSON(http.StatusOK, gin.H{"data": createdUser})
}

func createUser(user domain.RequestBody) entity.User{
	return entity.User{UserID: uuid.UUID.String(uuid.NewV4()),Name: user.Name}
}
