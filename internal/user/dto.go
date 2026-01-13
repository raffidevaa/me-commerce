package user

// reggister
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r RegisterRequest) ToUser() User {
	return User{
		Email:    r.Email,
		Password: r.Password,
	}
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func RegisterResponseFromUser(u User) RegisterResponse {
	return RegisterResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role,
	}
}

// login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
