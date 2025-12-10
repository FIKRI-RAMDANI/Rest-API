package service

import (
	"context"
	"errors"
	"time"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/FIKRI-RAMDANI/Rest-API/dto"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		conf:           cnf,
		userRepository: userRepository,
	}
}

func (a authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	// Server error in repository
	if err != nil {
		return dto.AuthResponse{}, errors.New("internal server Error")
	}
	// cek user
	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}
	// cek Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"sub": user.Id,
		"iss": "fikri-rest-api",
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("internal server error")
	}
	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}
