package domain

// struct untuk entity User
type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Balance  float64 `json:"balance"`
}

// struct untuk response yang dikembalikan data user yang login
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// struct untuk response yang dikembalikan ketika login
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// struct untuk melakukan login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// mendifinisikan method user repository dengan interface
type UserRepository interface {
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() ([]*User, error)
}

// mendifinisikan method user usecase dengan interface
type UserUsecase interface {
	GetUserBalance(id int) (*User, error)
	Login(login LoginRequest) (*LoginResponse, error)
	GetUserByID(id int) (*UserResponse, error)
}
