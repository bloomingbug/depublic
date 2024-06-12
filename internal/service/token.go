package service

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
)

type tokenService struct {
	otpRepository   repository.OneTimePasswordRepository
	tokenRepository repository.TokenRepository
}

func (t *tokenService) tokenGenerator() (token string) {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	c := make([]string, 64)
	for i := range c {
		numOrAlpha := rand.Intn(2)
		if numOrAlpha == 0 {
			c[i] = strconv.Itoa(randomizer.Intn(10))
		} else {
			c[i] = string(letters[randomizer.Intn(len(letters))])
		}

		token = strings.Join(c, "")
	}
	return
}

func (t *tokenService) GenerateTokenRegistration(c context.Context, otp, email string) (*entity.Token, error) {
	otpData, err := t.otpRepository.FindOneByCodeAndEmail(c, email, otp)
	if err != nil {
		return nil, err
	}

	err = t.otpRepository.Delete(c, otpData.ID)
	if err != nil {
		return nil, err
	}

	token := entity.NewTokenRegister(t.tokenGenerator(), email)
	token, err = t.tokenRepository.Create(c, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

type TokenService interface {
	GenerateTokenRegistration(c context.Context, otp, email string) (*entity.Token, error)
}

func NewTokenService(otpRepository repository.OneTimePasswordRepository, tokenRepository repository.TokenRepository) TokenService {
	return &tokenService{
		otpRepository:   otpRepository,
		tokenRepository: tokenRepository,
	}
}
