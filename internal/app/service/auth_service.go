package service

import (
	"github.com/api-voting/utils/exception"
	"github.com/api-voting/utils/security"
)

type AuthService interface {
	Login(username string, password string) (string, error)
	Logout(token string) error
}

type authService struct {
	service        UserService
	tokenBlacklist map[string]struct{}
}

func NewAuthService(service UserService) AuthService {
	return &authService{service: service, tokenBlacklist: make(map[string]struct{})}
}

func (s *authService) Login(username string, password string) (string, error) {
	user, err := s.service.FindByUsernamePassword(username, password)

	if err != nil {
		return "", err
	}

	token, err := security.CreateAccessToken(user)

	if err != nil {
		return "", exception.ErrFailedCreateToken
	}

	return token, nil
}

func (s *authService) Logout(token string) error {
	// Simpan token ke dalam blacklist
	s.tokenBlacklist[token] = struct{}{}
	return nil
}

func (s *authService) IsTokenBlacklisted(token string) bool {
	_, exists := s.tokenBlacklist[token]
	return exists
}
