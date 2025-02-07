package userdto

type RegisterRequest struct {
	Email    string `json:"identifier" example:"johndoe@example.com" minLength:"6" maxLength:"100"`
	Password string `json:"password" example:"xF54sal-M sa12" minLength:"8"`
}

type RegisterResponse struct{}
