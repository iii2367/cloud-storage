package jwt

import (
    "time"
	"cloud-storage/internal/config"
    "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Manager struct {
   	accessSecret  []byte
    refreshSecret []byte
    accessTTL  time.Duration
    refreshTTL time.Duration
    issuer  			string
}

func New(cfg config.JWT) *Manager {
    return &Manager {
        accessSecret:  	cfg.AccessSecret,
		refreshSecret: 	cfg.RefreshSecret,
        accessTTL:	 	cfg.AccessTTL,
		refreshTTL: 	cfg.RefreshTTL,
        issuer:  		cfg.Issuer,
    }
}

func (m *Manager) GenerateAccessToken(userID uint) (string, error) {
	now := time.Now()

    claims := AccessClaims {
		UserID:		userID, 
        RegisteredClaims: jwt.RegisteredClaims {
            Issuer:    m.issuer,
            ExpiresAt: jwt.NewNumericDate(now.Add(m.accessTTL)),
            IssuedAt:  jwt.NewNumericDate(now),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(m.accessSecret)
}

func (m *Manager) ParseAccessToken(tokenString string) (*AccessClaims, error) {
    claims := &AccessClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
        	return nil, ErrInvalidToken
	    }
        return m.accessSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, ErrInvalidToken
    }

    return claims, nil
}

func (m *Manager) GenerateRefreshToken(userID uint, sessionID uuid.UUID) (string, error) {
	now := time.Now()

	claims := RefreshClaims {
		UserID: userID,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims {
            Issuer:    m.issuer,
            ExpiresAt: jwt.NewNumericDate(now.Add(m.refreshTTL)),
            IssuedAt:  jwt.NewNumericDate(now),
        },
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(m.refreshSecret)
	if err != nil {
		return "", err
	}

    return tokenString, nil
}

func (m *Manager) ParseRefreshToken(tokenString string) (*RefreshClaims, error) {
	 claims := &RefreshClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
        	return nil, ErrInvalidToken
   	 	}
        return m.refreshSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, ErrInvalidToken
    }

    return claims, nil
}
