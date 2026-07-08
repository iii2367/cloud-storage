package config

import (
	"log"
	"time"
)

type JWT struct {
    Secret 	 []byte			
    Expires  time.Duration	
    Issuer   string			
}

func MustLoadJWT(prefix string) *JWT {
	expires, err := time.ParseDuration(mustEnv(prefix + "_JWT_EXPIRES"))
	if err != nil {
		log.Fatalf("invalid %s_JWT_EXPIRES: %v", prefix, err)
	}

	return &JWT {
		Secret:  []byte(mustEnv(prefix + "_JWT_SECRET")),
		Expires: expires,
		Issuer:  mustEnv(prefix + "_JWT_ISSUER"),
	}
}
