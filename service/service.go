package service

import (
	"golang.org/x/crypto/bcrypt"
	"user-management-api/domain"
	"user-management-api/entity"
	"user-management-api/repository"
)

//AddUserStatus - ENUM to represent status of AddUser function
type AddUserStatus string

const (
	//AddUserSuccess - 
	AddUserSuccess = "User added successfully"
	//UserAlreadyExists - 
	UserAlreadyExists = "User already exists"
	//InvalidUserID -
	InvalidUserID = "UserID format not valid"
	//InternalError -
	InternalError = "Internal Error"
)

//GetUser - Fetch user for given user ID from DB
func GetUser(userID string) (*entity.User, error) {
	user, err := userRepository.GetUser(userID)
	if(err != nil) {
		return nil, err
	}

	return user, nil
}

//AddUser - Add user to DB
func AddUser(requestBody domain.RequestBody) (AddUserStatus, error) {
	if (!isValidUserID(requestBody.UserID)) {
		return InvalidUserID, nil
	}
	
	createdUser, err := createUser(requestBody)

	if(err != nil) {
		return InternalError, err
	}
	
	err = userRepository.AddUser(*createdUser)

	if(err != nil) {
		return InternalError, err
	}

	return AddUserSuccess, nil
}

//AuthenticateUser - Authenticate given UserID and Password
func AuthenticateUser(userID string, password string) (bool, error) {
	user, err := userRepository.GetUser(userID)
	if(err != nil) {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if(err != nil) {
		return false, nil
	}

	return true, nil
}

func createUser(requestBody domain.RequestBody) (*entity.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if(err != nil) {
		return nil, err
	}

	return &entity.User{
		UserID: requestBody.UserID,
		Name: requestBody.Name,
		Password: string(password),
	}, nil
}

func isValidUserID(userID string) bool {
	if len(userID) == 0 {
		return false
	}

	return true
}

var userRepository repository.UserRepository

func init() {
	userRepository = repository.GetRepository()
}