package userdto

type LoginRequest struct {
	// Identifier could be any of username/phone/email and is at least 6 and
	// at most 100 characters long
	Identifier string `json:"identifier" example:"johndoe@example.com" minLength:"6" maxLength:"100"`
	Password   string `json:"password" example:"xF54sal-M sa12" minLength:"8"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
