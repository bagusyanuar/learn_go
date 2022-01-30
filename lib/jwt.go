package lib

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.StandardClaims
	Unique   uuid.UUID `json:"unique"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"roles"`
}

var ErrorBearerType = errors.New("invalid bearer")
var ErrorSignInMethod = errors.New("invalid signin method")
var ErrorJWTClaims = errors.New("invalid jwt claim")
var ErrorJWTParse = errors.New("invalid parse jwt")
var ErrorNoAuthorization = errors.New("invalid Unauthorized")

func CreateToken(unique uuid.UUID, username string, email string, role string) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: JWTIssuer,
		},
		Unique:   unique,
		Username: username,
		Email:    email,
		Role:     role,
	}
	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	signedToken, err := token.SignedString([]byte(JWTSignatureKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ClaimToken(auth string) (interface{}, error) {

	if auth == "" {
		return nil, ErrorNoAuthorization
	}
	bearer := string(auth[0:7])
	token := string(auth[7:])

	if bearer != "Bearer " {
		return nil, ErrorBearerType
	}

	vToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorSignInMethod
		} else if method != JWTSigninMethod {
			return nil, ErrorSignInMethod
		}
		return []byte(JWTSignatureKey), nil
	})
	if err != nil {
		return nil, ErrorJWTParse
	}
	claim, ok := vToken.Claims.(jwt.MapClaims)
	if !ok || !vToken.Valid {
		return nil, ErrorJWTClaims
	}
	return claim, nil
}
