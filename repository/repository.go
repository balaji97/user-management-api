package repository

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"user-management-api/entity"
)

const (
	awsRegion = "ap-south-1"
	awsCredentialsPath =  "/Users/balaji/.aws/credentials"
	awsProfile = "user-management-api"
	tableName = "users"
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
func (dynamoDBRepository *DynamoDBRepository) GetUser(UserID string) (*entity.User, error) {
	getItemInput, err := buildDynamoDBGetInput(entity.UserKey{UserID: UserID})
	if(err != nil) {
		return nil, err
	}

	user := entity.User{}
	getResult, err := dynamoDBRepository.repository.GetItem(getItemInput)
	if(err != nil) {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalMap(getResult.Item, &user)
	if(err != nil) {
		return nil, err
	}

	return &user, nil
}

//AddUser - Add user to repository
func (dynamoDBRepository *DynamoDBRepository) AddUser(User entity.User) error{
	putItemInput, err := buildDynamoDBPutInput(User)
	if(err != nil) {
		return err
	}

	_, err = dynamoDBRepository.repository.PutItem(putItemInput)

	return err
}

func buildDynamoDBGetInput(userKey entity.UserKey) (*dynamodb.GetItemInput, error){
	key, err := dynamodbattribute.MarshalMap(userKey)
	if(err != nil) {
		return nil, err
	}

	return &dynamodb.GetItemInput{Key: key, TableName: aws.String(tableName)}, nil
}

func buildDynamoDBPutInput(user entity.User) (*dynamodb.PutItemInput, error) {
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, err
	}

	return &dynamodb.PutItemInput{Item: item, TableName: aws.String(tableName)}, nil
}

var repository *DynamoDBRepository

//InitializeRepository - Set up connection to DynamoDB repository
func InitializeRepository() error{
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewSharedCredentials(awsCredentialsPath, awsProfile)})

	if(err != nil) {
		return err
	}

	repository = &DynamoDBRepository{
		repository: dynamodb.New(session)} 
	
	return nil
}

//GetRepository - Returns an implementation of UserRepository
func GetRepository() *DynamoDBRepository {
	return repository
}