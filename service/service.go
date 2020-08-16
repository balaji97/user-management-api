package service

import (
	"golang.org/x/crypto/bcrypt"
	"user-management-api/domain"
	"user-management-api/entity"
	"user-management-api/repository"
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
func AddUser(requestBody domain.RequestBody) (domain.AddUserStatus, error) {
	if (!isValidUserID(requestBody.UserID)) {
		return domain.InvalidUserID, nil
	}
	
	createdUser, err := createUser(requestBody)

	if(err != nil) {
		return domain.InternalError, err
	}
	
	err = userRepository.AddUser(*createdUser)

	if(err != nil) {
		if(err.Error() == domain.UserAlreadyExists) {
			return domain.UserAlreadyExists, err
		}

		return domain.InternalError, err
	}

	return domain.AddUserSuccess, nil
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