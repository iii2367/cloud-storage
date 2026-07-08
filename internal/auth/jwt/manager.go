package jwt

import (
    "errors"
    "time"
	"cloud-storage/internal/config"
    "github.com/golang-jwt/jwt/v5"
)

type Manager struct {
    secret  []byte
    expires time.Duration
    issuer  string
}

func New(cfg config.JWT) *Manager {
    return &Manager {
        secret:  cfg.Secret,
        expires: cfg.Expires,
        issuer:  cfg.Issuer,
    }
}

func (m *Manager) Generate(id uint) (string, error) {
    claims := Claims {
		ID:		id, 
        RegisteredClaims: jwt.RegisteredClaims {
            Issuer:    m.issuer,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.expires)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(m.secret)
}

func (m *Manager) Parse(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return m.secret, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
