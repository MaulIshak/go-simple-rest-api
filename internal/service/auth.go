package service

import (
	"context"
	"errors"
	"go-simple-rest-api/domain"
	"go-simple-rest-api/dto"
	"go-simple-rest-api/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           config.Config
	userRepository domain.UserRepository
}


func NewAuth(cnf config.Config, userRpository domain.UserRepository) domain.AuthService {
	return authService{
		conf:           cnf,
		userRepository: userRpository,
	}
}

// Login implements domain.AuthService.
func (a authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil{
		return dto.AuthResponse{}, err
	}
	if user.Id == ""{
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil{
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	claim := jwt.MapClaims{
		"id":user.Id,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))

	if err != nil{
		return dto.AuthResponse{}, errors.New("authentication gagal")
	}

	return  dto.AuthResponse{
		Token: tokenStr,

	}, nil
}