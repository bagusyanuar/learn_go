package auth

import (
	"net/http"
	"web_go/lib"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth(context *gin.Context) {
	auth := context.Request.Header.Get("Authorization")

	claim, err := lib.ClaimToken(auth)

	if err != nil {
		var message string = ""

		switch err {
		case lib.ErrorBearerType:
			message = "Invalid Bearer Type"
		case lib.ErrorSignInMethod:
			message = "Invalid SignIn Method"
		case lib.ErrorJWTParse:
			message = "Invalid Token Parse"
		case lib.ErrorJWTClaims:
			message = "Invalid Token Claim"
		default:
			message = "Unauthenticated"
		}
		context.AbortWithStatusJSON(http.StatusUnauthorized, lib.ApiResponse{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: message,
		})
		return
	}
	context.Set("user", claim)
	context.Next()
}

func IsMember(context *gin.Context) {
	user := context.MustGet("user").(jwt.MapClaims)
	if user["roles"] != "member" {
		context.AbortWithStatusJSON(http.StatusForbidden, lib.ApiResponse{
			Code:    http.StatusForbidden,
			Data:    nil,
			Message: "Forbidden Access",
		})
		return
	}
	context.Next()
}
