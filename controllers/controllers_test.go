package controllers

import (
	"testing"
	"user-management-api/domain"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	requestBody := domain.RequestBody{Name: "someName"}
	user := createUser(requestBody)

	assert.Equal(t, requestBody.Name, user.Name)
	assert.NotEmpty(t, user.UserID)
	
}
