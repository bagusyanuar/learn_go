package controller

import (
	"errors"
	"net/http"
	"web_go/database"
	"web_go/lib"
	"web_go/model"
	"web_go/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMemberProfile(context *gin.Context) {
	user := context.MustGet("user").(jwt.MapClaims)
	var member model.UserMember
	err := database.Db.Debug().Preload("Member").Joins("JOIN member ON users.id = member.user_id").Where("users.id = ?", user["unique"]).First(&member).Error
	if err != nil {
		eMsg := "Server Failure"
		if errors.Is(err, gorm.ErrRecordNotFound) {
			eMsg = "User Not Found"
		}
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    500,
			Data:    nil,
			Message: eMsg,
		})
		return
	}

	context.JSON(http.StatusOK, lib.ApiResponse{
		Code:    200,
		Data:    member,
		Message: "Succes Get Profile",
	})
}
func GetMembers(ctx *gin.Context) {
	var members []model.Member
	err := repository.FindAllMembers(&members)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, members)
	}
}

func GetMemberUsers(context *gin.Context) {
	// var member model.MemberUser
	// var apiModel model.MemberUser
	var api []model.ResponseMemberUser
	var err error
	if err = database.Db.Debug().
		Select(
			"`member`.`id`",
			"`member`.`full_name`",
			"`member`.`phone`",
			"`member`.`school`",
			"`member`.`user_id`",
			"`User`.`id` as `User__id`",
			"`User`.`email` `User___email`").
		Joins("INNER JOIN `users` `User` ON `member`.`user_id` = `User`.`id`").Find(&api).Error; err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// if err = database.Db.Debug().Table("member").Preload("User").Find(&apiModel).Error; err != nil {
	// 	context.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }
	// return nil
	// err := repository.FindAllMembersWithUser(&memberUser, &model.MemberUser{})
	// if err != nil {
	// 	context.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }
	context.JSON(http.StatusOK, api)
}
