package domain

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

//AddUserStatus - ENUM to represent status of AddUser function
type AddUserStatus string

//RequestBody - Domain object to parse HTTP JSON request body
type RequestBody struct {
	UserID, Name, Password string
}