package dto

// LoginRequest represents the login credentials
type LoginRequest struct {
	// Username must start with a letter and be 6-31 characters long
	// @pattern ^[a-zA-Z]{1}[a-zA-Z0-9]{5,30}$
	Username string `json:"username" example:"johndoe"`

	// User's password
	// @minLength 8
	Password string `json:"password" example:"********"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	// JWT token for authentication
	// @example eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	Token string `json:"token"`
}
