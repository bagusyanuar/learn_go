package controller

import (
	"encoding/json"
	"net/http"
	"web_go/lib"
	"web_go/model"
	"web_go/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(ctx *gin.Context) {
	var users []model.User
	err := repository.FindAllUsers(&users)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

type createRequest struct {
	Email    *string
	Password string
	Provider *string
	Name     *string
}

func CreateUserMember(context *gin.Context) {
	var request createRequest
	if err := context.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if tempEmail := request.Email; tempEmail == nil {
		context.JSON(http.StatusBadRequest, lib.ApiResponse{Code: http.StatusBadRequest, Data: nil, Message: "Email Cannot Be empty"})
		return
	}

	if tempName := request.Name; tempName == nil {
		context.JSON(http.StatusBadRequest, lib.ApiResponse{Code: http.StatusBadRequest, Data: nil, Message: "Name Cannot Be Empty"})
		return
	}

	if prov := request.Provider; prov == nil {
		context.JSON(http.StatusBadRequest, lib.ApiResponse{Code: http.StatusBadRequest, Data: nil, Message: "Provider Cannot Be Null"})
		return
	}

	var email string = *request.Email
	var name string = *request.Name
	roles, _ := json.Marshal([]string{"ROLE_MEMBER"})
	provider, _ := json.Marshal([]string{*request.Provider})
	password := []byte(request.Password)

	hashedPasswod, errorHashed := bcrypt.GenerateFromPassword(password, 13)
	if errorHashed != nil {
		panic(errorHashed.Error())
	}
	var vPassword string = string(hashedPasswod)
	user := model.User{
		Password:   vPassword,
		Email:      email,
		IsVerified: true,
		Roles:      roles,
		Provider:   provider,
	}

	member := model.Member{
		FullName: name,
	}
	userMember := model.UserMember{
		User:   user,
		Member: member,
	}
	err := repository.CreateUserMember(&userMember)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	response := lib.ApiResponse{Code: 200, Data: user, Message: "oke"}
	context.JSON(http.StatusOK, response)

}

func CreateSingularUser(context *gin.Context) {
	var request createRequest
	var email string = *request.Email
	roles, _ := json.Marshal([]string{"ROLE_MEMBER"})
	provider, _ := json.Marshal([]string{*request.Provider})
	password := []byte(request.Password)

	hashedPasswod, errorHashed := bcrypt.GenerateFromPassword(password, 13)
	if errorHashed != nil {
		panic(errorHashed.Error())
	}
	var vPassword string = string(hashedPasswod)
	user := model.User{
		Password:   vPassword,
		Email:      email,
		IsVerified: true,
		Roles:      roles,
		Provider:   provider,
	}
	err := repository.CreateUser(&user)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	response := lib.ApiResponse{Code: 200, Data: user, Message: "oke"}
	context.JSON(http.StatusOK, response)

}
