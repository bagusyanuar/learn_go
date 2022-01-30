package lib

import "github.com/dgrijalva/jwt-go"

var JWTSignatureKey string
var JWTIssuer string
var JWTSigninMethod = jwt.SigningMethodHS256
