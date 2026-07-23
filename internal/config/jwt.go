package config

import (
	"log"
	"time"
)

type JWT struct {
	AccessSecret  []byte
	RefreshSecret []byte
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
	Issuer        string
}

func MustLoadJWT(prefix string) *JWT {
	accessTTL, err := time.ParseDuration(mustEnv(prefix + "_JWT_ACCESSTTL"))
	if err != nil {
		log.Fatalf("invalid %s_JWT_ACCESSTTL: %v", prefix, err)
	}
	refreshTTL, err := time.ParseDuration(mustEnv(prefix + "_JWT_REFRESHTTL"))
	if err != nil {
		log.Fatalf("invalid %s_JWT_REFRESHTTL: %v", prefix, err)
	}

	return &JWT{
		AccessSecret:  []byte(mustEnv(prefix + "_JWT_ACCESS_SECRET")),
		RefreshSecret: []byte(mustEnv(prefix + "_JWT_REFRESH_SECRET")),
		AccessTTL:     accessTTL,
		RefreshTTL:    refreshTTL,
		Issuer:        mustEnv(prefix + "_JWT_ISSUER"),
	}
}
