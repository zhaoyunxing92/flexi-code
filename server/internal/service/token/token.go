package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/zhaoyunxing92/flexi-code/server/configs"
	"github.com/zhaoyunxing92/flexi-code/server/internal/domain"
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/sys"
)

type Service struct {
	secret  string
	expired uint
}

func NewTokenService(config *configs.App) *Service {
	return &Service{secret: config.Secret, expired: config.Expired}
}

func (s *Service) GenerateToken(account sys.Account) (string, error) {
	expired := time.Now().Add(time.Duration(s.expired) * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.NewClaims(account, expired))
	return token.SignedString([]byte(s.secret))
}

func (s *Service) ParseToken(code string) (*domain.FlexiCodeClaims, error) {
	token, err := jwt.ParseWithClaims(code, &domain.FlexiCodeClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*domain.FlexiCodeClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
