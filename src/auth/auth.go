package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"web_go/database"
	"web_go/lib"
	"web_go/model"
	"web_go/repository"
	req "web_go/src/request"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminSignUp(context *gin.Context) {
	var request req.AdminSignRequest
	context.Bind(&request)
	roles, _ := json.Marshal([]string{"ROLE_ADMIN"})
	var requiredField = []string{request.Email, request.Password}

	isValid := lib.ValidateEmptyRequest(requiredField...)
	if !isValid {
		context.JSON(http.StatusBadRequest, lib.ApiResponse{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: "Bad Request! Parameter cannot be empty",
		})
		return
	}
	provider, _ := json.Marshal([]string{"app"})
	password := ""
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
	if errHashing != nil {
		context.JSON(http.StatusBadRequest, lib.ApiResponse{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: "Bad Request! Failed Hashing Password",
		})
		return
	}
	password = string(hash)

	// insert using transaction when tokenize jwt
	tx := database.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	admin := model.User{
		Email:      request.Email,
		Password:   password,
		Provider:   provider,
		Roles:      roles,
		IsVerified: true,
	}

	if err := tx.Create(&admin).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error Insert",
		})
		return
	}

	accessToken, errTokenize := lib.CreateToken(admin.Id, "Administrator", admin.Email, "admin")
	if errTokenize != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error while creating token",
		})
		return
	}

	tx.Commit()
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code: 200,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign Up",
	})
}

func AdminSignIn(context *gin.Context) {
	var request req.AdminSignRequest
	context.Bind(&request)
	var admin model.User
	clause := fmt.Sprintf("JSON_SEARCH(%s, 'all', '%s') IS NOT NULL", "roles", "ROLE_ADMIN")
	err := database.Db.Debug().Where(clause).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusUnauthorized, lib.ApiResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: "User Not Found!",
			})
			return
		} else if errors.Is(err, lib.ErrorInvalidPassword) {
			context.JSON(http.StatusUnauthorized, lib.ApiResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: "Password Did Not Match!",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    500,
			Data:    nil,
			Message: "Error While Signin " + err.Error(),
		})
		return
	}

	accessToken, err := lib.CreateToken(admin.Id, "Administrator", admin.Email, "admin")
	if err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    500,
			Data:    nil,
			Message: "Error While Getting Access Token " + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code: 200,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign In",
	})
}

func MemberSignUp(context *gin.Context) {
	var request req.MemberSignUpRequest
	context.Bind(&request)
	roles, _ := json.Marshal([]string{"ROLE_MEMBER"})
	var requiredField = []string{request.Email, request.Name, request.Provider}
	if request.Provider == "app" {
		requiredField = append(requiredField, request.Password)

	}
	isValid := lib.ValidateEmptyRequest(requiredField...)
	if !isValid {
		context.JSON(http.StatusBadRequest, lib.ApiResponse{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: "Bad Request! Parameter cannot be empty",
		})
		return
	}
	provider, _ := json.Marshal([]string{request.Provider})
	password := ""
	if request.Provider == "app" {
		hash, errHashing := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
		if errHashing != nil {
			context.JSON(http.StatusBadRequest, lib.ApiResponse{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "Bad Request! Failed Hashing Password",
			})
			return
		}
		password = string(hash)
	}

	// insert using transaction when tokenize jwt
	tx := database.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	member := model.UserMember{
		User: model.User{
			Email:      request.Email,
			Password:   password,
			Provider:   provider,
			Roles:      roles,
			IsVerified: true,
		},
		Member: model.Member{
			FullName: request.Name,
			School:   request.School,
			Avatar:   request.Avatar,
			Address:  request.Address,
			Phone:    request.Phone,
		},
	}

	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error Insert",
		})
		return
	}

	accessToken, errTokenize := lib.CreateToken(member.Id, member.Member.FullName, member.Email, "member")
	if errTokenize != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error while creating token",
		})
		return
	}

	tx.Commit()
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code: 200,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign Up",
	})
}

func MemberSignIn(context *gin.Context) {
	var request req.MemberSignInRequest
	context.Bind(&request)
	user, err := repository.SignInMember(&model.UserMember{}, request.Email, request.Password, request.Provider)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusUnauthorized, lib.ApiResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: "User Not Found!",
			})
			return
		} else if errors.Is(err, lib.ErrorInvalidPassword) {
			context.JSON(http.StatusUnauthorized, lib.ApiResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: "Password Did Not Match!",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    500,
			Data:    nil,
			Message: "Error While Signin " + err.Error(),
		})
		return
	}

	accessToken, err := lib.CreateToken(user.Id, user.Member.FullName, user.Email, "member")
	if err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    500,
			Data:    nil,
			Message: "Error While Getting Access Token " + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code: 200,
		Data: map[string]interface{}{
			"accessToken": accessToken,
		},
		Message: "Success Sign In",
	})
}

// Func Sign Up
// var request signUpRequest
// 	context.Bind(&request)
// 	roles, _ := json.Marshal([]string{"ROLE_MEMBER"})
// 	request.Roles = roles
// 	isValid := lib.ValidateEmptyRequest(request.Email, request.Name, request.Provider, request.Password)
// 	if !isValid {
// 		context.JSON(http.StatusBadRequest, lib.ApiResponse{
// 			Code:    http.StatusBadRequest,
// 			Data:    nil,
// 			Message: "Bad Request! Parameter cannot be empty",
// 		})
// 		return
// 	}
// 	provider, _ := json.Marshal([]string{request.Provider})
// 	hashedPasswod, errHashed := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
// 	if errHashed != nil {
// 		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
// 			Code:    http.StatusInternalServerError,
// 			Data:    nil,
// 			Message: "Internal Server Error! Error while hashing password",
// 		})
// 		return
// 	}

// 	var vPassword string = string(hashedPasswod)
// 	user := model.UserMember{
// 		User: model.User{
// 			Email:      request.Email,
// 			Password:   vPassword,
// 			Provider:   provider,
// 			Roles:      request.Roles,
// 			IsVerified: true,
// 		},
// 		Member: model.Member{
// 			FullName: request.Name,
// 		},
// 	}

// 	err := repository.RegisterMember(&user)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
// 			Code:    http.StatusInternalServerError,
// 			Data:    nil,
// 			Message: "Internal Server Error! Error while register user " + err.Error(),
// 		})
// 		return
// 	}

// 	claims := JWTClaims{
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer: JWT_ISSUER,
// 		},
// 		Username: user.Member.FullName,
// 		Email:    user.Email,
// 		Roles:    "member",
// 	}

// 	token := jwt.NewWithClaims(JWT_SIGNIN_METHOD, claims)
// 	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
// 			Code:    500,
// 			Data:    nil,
// 			Message: "Error While Signin JWT " + err.Error(),
// 		})
// 		return
// 	}
// 	context.JSON(http.StatusOK, lib.ApiResponse{
// 		Code: 200,
// 		Data: map[string]interface{}{
// 			"token": signedToken,
// 		},
// 		Message: "Success Sign Up",
// 	})

// func ClaimJWT(context *gin.Context) {
// 	vToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJKT05JIiwidXNlcm5hbWUiOiJqb25pIGVzbW9kIiwiZW1haWwiOiJqb25pQGdtYWlsLmNvbSJ9.SQo6n-14VDe-PGEvs9D93GR756ahBfzNwz0UvFx8g5s"

// 	token, err := jwt.Parse(vToken, func(t *jwt.Token) (interface{}, error) {
// 		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("signin method invalid")
// 		} else if method != JWT_SIGNIN_METHOD {
// 			return nil, fmt.Errorf("signing method invalid")
// 		}
// 		return JWT_SIGNATURE_KEY, nil
// 	})
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
// 			Code:    500,
// 			Data:    nil,
// 			Message: "Error Parse " + err.Error(),
// 		})
// 		return
// 	}

// 	claim, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
// 			Code:    500,
// 			Data:    nil,
// 			Message: "Error Claim " + err.Error(),
// 		})
// 		return
// 	}

// 	context.JSON(http.StatusOK, lib.ApiResponse{
// 		Code:    200,
// 		Data:    claim,
// 		Message: "Success Sign Up",
// 	})
// }
