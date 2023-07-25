package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/sys"
)

type FlexiCodeClaims struct {
	Account string `json:"account"`
	// account avatar
	Avatar string `json:"avatar"`

	Name string `json:"name"`

	EffectiveTime  time.Time `json:"-"`
	ExpirationTime time.Time `json:"-"`
	CreationTime   time.Time `json:"-"`
}

func NewClaims(account sys.Account, expired time.Time) *FlexiCodeClaims {
	return &FlexiCodeClaims{
		Account:        account.AccountId.String,
		Avatar:         account.Avatar,
		Name:           account.Name,
		ExpirationTime: expired,
		EffectiveTime:  time.Now(),
		CreationTime:   time.Now(),
	}
}

func (c *FlexiCodeClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c *FlexiCodeClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(c.CreationTime), nil
}

func (c *FlexiCodeClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(c.EffectiveTime), nil
}
func (c *FlexiCodeClaims) GetIssuer() (string, error) {
	return "fc", nil
}

func (c *FlexiCodeClaims) GetSubject() (string, error) {
	return "fc", nil
}

func (c *FlexiCodeClaims) GetAudience() (jwt.ClaimStrings, error) {
	return []string{"fc"}, nil
}
