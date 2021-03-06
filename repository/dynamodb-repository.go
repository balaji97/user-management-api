package repository

import (
	"errors"
	"user-management-api/domain"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"user-management-api/entity"
)

const (
	awsRegion = "ap-south-1"
	awsCredentialsPath =  "/Users/balaji/.aws/credentials"
	awsProfile = "user-management-api"
	tableName = "users"
	putConditionalExpression = "attribute_not_exists(UserID)"
	conditionalFailed = "ConditionalCheckFailedException: The conditional request failed"
)

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
	} else if(getResult.Item == nil) {
		return nil, nil
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

	if(err != nil) {
		if(err.Error() == conditionalFailed) {
			return errors.New(domain.UserAlreadyExists)
		}
	}

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

	conditionalExpression := putConditionalExpression

	return &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(tableName),
		ConditionExpression: &conditionalExpression,
		}, nil
}

func init() {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewSharedCredentials(awsCredentialsPath, awsProfile)})

	if(err != nil) {
		panic(err)
	}

	dynamoDBRepository = &DynamoDBRepository{repository: dynamodb.New(session)} 
}

var dynamoDBRepository *DynamoDBRepository