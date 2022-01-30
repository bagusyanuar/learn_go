package controller

import (
	"errors"
	"net/http"
	"path/filepath"
	"web_go/database"
	"web_go/lib"
	"web_go/model"
	req "web_go/src/request"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CategoryCreate(context *gin.Context) {
	var request req.CategoryRequest
	context.Bind(&request)

	var iconName *string

	if request.Icon != nil {
		ext := filepath.Ext(request.Icon.Filename)
		filename := "assets/icons/" + uuid.New().String() + ext
		iconName = &filename
		if errUpload := context.SaveUploadedFile(request.Icon, filename); errUpload != nil {
			context.JSON(http.StatusInternalServerError, lib.ApiResponse{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "Internal Server Error. Failed while upload icon",
			})
			return
		}
	}

	category := model.Categories{
		Name: request.Name,
		Slug: lib.MakeSlug(request.Name),
		Icon: iconName,
	}
	if err := database.Db.Debug().Create(&category).Error; err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error. Failed While Insert",
		})
		return
	}
	response := lib.ApiResponse{Code: 200, Data: category, Message: "Success"}
	context.JSON(http.StatusOK, response)
}

func GetCategories(context *gin.Context) {
	param := context.Query("name")
	var categories []model.Categories
	err := database.Db.Debug().Where("name LIKE ?", "%"+param+"%").Find(&categories).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error..",
		})
		return
	}
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code:    200,
		Data:    categories,
		Message: "Success",
	})
}

func GetDetailCategories(context *gin.Context) {
	param := context.Param("id")
	var category model.Categories
	if err := database.Db.Debug().First(&category, param).Error; err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error..",
		})
		return
	}
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code:    200,
		Data:    category,
		Message: "Success",
	})
}
func PatchCategories(context *gin.Context) {
	param := context.Param("id")
	var request req.CategoryRequest
	context.Bind(&request)

	var category model.Categories
	errFind := database.Db.Debug().Where("id = ?", param).First(&category).Error

	if errFind != nil {
		if errors.Is(errFind, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusInternalServerError, lib.ApiResponse{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "User Not Found!",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Error While Find Categories ",
		})
		return
	}

	category.Name = request.Name
	category.Slug = lib.MakeSlug(request.Name)
	if err := database.Db.Debug().Save(&category).Error; err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error. Failed While Update",
		})
		return
	}
	response := lib.ApiResponse{Code: 200, Data: category, Message: "Success"}
	context.JSON(http.StatusOK, response)
}

func DeleteCategories(context *gin.Context) {
	param := context.Param("id")
	var category model.Categories
	err := database.Db.Debug().Delete(&category, param).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, lib.ApiResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error..",
		})
		return
	}
	context.JSON(http.StatusOK, lib.ApiResponse{
		Code:    200,
		Data:    nil,
		Message: "Success",
	})
}
