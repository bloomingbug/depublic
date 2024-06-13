package service

import (
	"context"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
)

type tokenService struct {
	otpRepository   repository.OneTimePasswordRepository
	tokenRepository repository.TokenRepository
}

func (t *tokenService) GenerateTokenRegistration(c context.Context, otp, email string) (*entity.Token, error) {
	otpData, err := t.otpRepository.FindOneByCodeAndEmail(c, email, otp)
	if err != nil || otpData == nil {
		return nil, err
	}

	err = t.otpRepository.Delete(c, otpData.ID)
	if err != nil {
		return nil, err
	}

	token := entity.NewToken(email, entity.Register)
	token, err = t.tokenRepository.Create(c, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

type TokenService interface {
	GenerateTokenRegistration(c context.Context, otp, email string) (*entity.Token, error)
}

func NewTokenService(
	otpRepository repository.OneTimePasswordRepository,
	tokenRepository repository.TokenRepository) TokenService {
	return &tokenService{
		otpRepository:   otpRepository,
		tokenRepository: tokenRepository,
	}
}
