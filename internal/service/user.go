package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	tokenRepository repository.TokenRepository
	userRepository  repository.UserRepository
	jwtToken        jwt_token.JwtToken
}

func (s *userService) UserRegistration(
	c echo.Context,
	token string,
	email string,
	user *entity.User) (*entity.User, error) {
	tokenData, err := s.tokenRepository.FindOneByCodeAndEmail(c.Request().Context(), user.Email, token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if tokenData.Action != "register" {
		return nil, errors.New("invalid token")
	}

	err = s.tokenRepository.Delete(c.Request().Context(), tokenData.ID)
	if err != nil {
		return nil, err
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return nil, err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	avatarName := fmt.Sprintf("%s-%s", strconv.FormatInt(time.Now().Unix(), 10), file.Filename)
	fullPath := filepath.Join(currentDir, "storage", "user", "avatars", avatarName)

	err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.Avatar = avatarName

	user, err = s.userRepository.Create(c.Request().Context(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(c context.Context, email, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(c, email)
	if err != nil {
		return "", errors.New("email/password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email/password is incorrect")
	}

	now := time.Now()
	expiredTime := now.Local().Add(24 * time.Hour)
	claims := jwt_token.JwtCustomClaims{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Role:   string(user.Role),
		Avatar: user.Avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Depublic",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.jwtToken.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("kesalahan sistem")
	}

	return token, nil
}

func (s *userService) FindUserByEmail(c context.Context, email string) (*entity.User, error) {
	user, err := s.userRepository.FindByEmail(c, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type UserService interface {
	UserRegistration(c echo.Context, token string, email string, user *entity.User) (*entity.User, error)
	Login(c context.Context, email, password string) (string, error)
	FindUserByEmail(c context.Context, email string) (*entity.User, error)
}

func NewUserService(
	tokenRepository repository.TokenRepository,
	userRepository repository.UserRepository,
	jwtToken jwt_token.JwtToken) UserService {
	return &userService{
		tokenRepository: tokenRepository,
		userRepository:  userRepository,
		jwtToken:        jwtToken,
	}
}
