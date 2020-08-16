package service

import (
	"testing"
	"user-management-api/domain"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	requestBody := domain.RequestBody{Name: "someName"}
	user, err := createUser(requestBody)

	assert.Nil(t, err)
	assert.Equal(t, requestBody.Name, user.Name)
	assert.NotEmpty(t, user.UserID)
	assert.NotEmpty(t, user.Password)
	
}
