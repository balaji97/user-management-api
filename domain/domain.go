package domain

//RequestBody - Domain object to parse HTTP JSON request body
type RequestBody struct {
	UserID, Name, Password string
}