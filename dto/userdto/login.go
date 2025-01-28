package userdto

type LoginRequest struct {
	// Username must start with a letter and have a length of 6-31 characters
	Username string `json:"username" example:"john doe" minLength:"6" maxLength:"31"`
	Password string `json:"password" example:"xF54sal-M" minLength:"8"`
}

type LoginResponse struct {
	// Token JWT token for authentication
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
