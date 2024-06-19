package jwt

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
)

type selfClaims struct {
	Payload []byte `json:"payload"`
	jwt.RegisteredClaims
}

func newClaims(claims IClaims) (*selfClaims, error) {
	payload, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}
	c := &selfClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    claims.GetIssuer(),
			Subject:   claims.GetSubject(),
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(claims.GetExpirationTime()),
			NotBefore: jwt.NewNumericDate(claims.GetIssuedAt()),
			IssuedAt:  jwt.NewNumericDate(claims.GetIssuedAt()),
			ID:        claims.GetSubject(),
		},
	}
	return c, nil
}
