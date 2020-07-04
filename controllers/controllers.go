package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController - "/"
func HomeController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
}
