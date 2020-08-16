package service

import (
	"github.com/satori/go.uuid"
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
func AddUser(requestBody domain.RequestBody) error {
	createdUser, err := createUser(requestBody)

	if(err != nil) {
		return err
	}
	
	err = userRepository.AddUser(*createdUser)

	if(err != nil) {
		return err
	}

	return nil
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
		UserID: uuid.UUID.String(uuid.NewV4()),
		Name: requestBody.Name,
		Password: string(password),
	}, nil
}

var userRepository repository.UserRepository

func init() {
	userRepository = repository.GetRepository()
}