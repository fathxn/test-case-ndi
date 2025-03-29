package usecase

import (
	"errors"
	"test-case-ndi/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// struct untuk userUsecase
type userUsecase struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

// struct untuk jwt claims
type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// constructor untuk userUsecase
func NewUserUsecase(userRepo domain.UserRepository, jwtSecret string) domain.UserUsecase {
	return &userUsecase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// implementasi method GetUserBalance
func (u *userUsecase) GetUserBalance(id int) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

// implementasi method Login
func (u *userUsecase) Login(creds domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := u.userRepo.GetByUsername(creds.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Password != creds.Password {
		return nil, errors.New("invalid credentials")
	}

	claims := JWTClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	userResponse := domain.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return &domain.LoginResponse{
		Token: tokenString,
		User:  userResponse,
	}, nil
}

// implementasi method GetUserByID
func (u *userUsecase) GetUserByID(id int) (*domain.UserResponse, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &domain.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
