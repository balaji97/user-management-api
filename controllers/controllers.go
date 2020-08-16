package controllers

import (
	"user-management-api/service"
	"user-management-api/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)


// Home - GET("/")
func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
}

// GetUser - GET("/user/{UserID}") - Return User as JSON object for given ID
func GetUser(context *gin.Context) {
	user, err := service.GetUser(context.Param("UserID"))

	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else if(user == nil) {
		context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": user})
}

//AddUser - POST("/user") - Add user as per POST body
func AddUser(context *gin.Context) {
	var requestBody domain.RequestBody
	context.BindJSON(&requestBody)

	err := service.AddUser(requestBody)
	
	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "OK"})
}

//AuthenticateUser - Checks if given User ID and password are valid
func AuthenticateUser(context *gin.Context) {

	authenticationStatus, err := service.AuthenticateUser(context.Param("UserID"), context.Param("Password"))
	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	if(authenticationStatus == false) {
		context.JSON(http.StatusUnauthorized, gin.H{"data": "Authentication failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "OK"})
}
