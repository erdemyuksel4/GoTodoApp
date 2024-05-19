package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &JWTClaim{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	claim, err := TokenToClaim(signedToken)
	if err != nil {
		return err
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}

	return

}
func TokenToClaim(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	return claims, nil
}
