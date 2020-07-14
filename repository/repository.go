package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"user-management-api/entity"
)

//UserRepository - Interface defining repository calls related to users database
type UserRepository interface {
	AddUser(User entity.User) error
	GetUser(UserID string) (entity.User, error)
}

//DynamoDBRepository - Implementation of UserRepository
type DynamoDBRepository struct {
	repository *dynamodb.DynamoDB
}

//GetUser - Return user for given UserId
func (dynamoDBRepository *DynamoDBRepository) GetUser(UserID string) entity.User {
	return entity.User{UserID: UserID, Name: "name"}
}

//AddUser - Add user to repository
func (dynamoDBRepository *DynamoDBRepository) AddUser(User entity.User) {
	//Do nothing
}

//GetRepository - Returns an implementation of UserRepository
func GetRepository() *DynamoDBRepository {
	if(repository == nil) {
		//TODO - Initialize repository here
		return repository
	}
			
	return repository	
}

var repository *DynamoDBRepository