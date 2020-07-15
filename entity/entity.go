package entity

//User - Entity object to store User details
type User struct {
	UserID, Name string
}

//UserKey - Object to pass DynamoDB key of User
type UserKey struct {
	UserID string
}
