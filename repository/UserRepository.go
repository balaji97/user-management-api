package repository

import (
	"user-management-api/entity"
)

//GetUser - Return user for given UserId
func GetUser(UserID string) entity.User{
	return entity.User{UserID: UserID, Name: "name"}
}