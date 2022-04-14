package injection

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
	leafJwt "github.com/paulusrobin/leaf-utilities/token/jwt"
)

func NewToken(tokenConfig pkgConfig.TokenConfig) (leafToken.Token, error) {
	return leafJwt.NewJWT(tokenConfig.TokenPublicKey, tokenConfig.TokenAlg), nil
}
