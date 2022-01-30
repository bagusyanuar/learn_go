package controller

import (
	"net/http"
	"web_go/lib"
	"web_go/model"
	"web_go/repository"

	"github.com/gin-gonic/gin"
)

func GetTeachers(context *gin.Context) {
	var teachers []model.Teacher
	err := repository.FindAllTeachers(&teachers)
	if err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error",
		})
		return
	}
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code:    http.StatusOK,
		Data:    teachers,
		Message: "Success",
	})
}
