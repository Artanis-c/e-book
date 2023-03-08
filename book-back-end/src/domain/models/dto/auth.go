package result

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	OpenId string
	UserNo string
	jwt.StandardClaims
}
