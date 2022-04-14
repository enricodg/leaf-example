package pkgConfig

import leafConfig "github.com/paulusrobin/leaf-utilities/config"

type TokenConfig struct {
	// - Token Settings
	TokenPublicKey string `envconfig:"TOKEN_PUBLIC_KEY" default:"ZYxEVwK3tmj3hXLV" required:"true"`
	TokenAlg       string `envconfig:"TOKEN_ALG" required:"true" default:"HS256"`
}

func NewTokenConfig() (TokenConfig, error) {
	configuration := TokenConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return TokenConfig{}, err
	}
	return configuration, nil
}
