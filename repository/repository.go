package repository

import (
	"user-management-api/entity"
)

//UserRepository - Interface defining repository calls related to users database
type UserRepository interface {
	AddUser(User entity.User) error
	GetUser(UserID string) (*entity.User, error)
}

//GetRepository - Returns an implementation of UserRepository
func GetRepository() *DynamoDBRepository {
	return dynamoDBRepository
}