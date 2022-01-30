package lib

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

func IsPasswordValid(plainPassword string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	return err == nil
}
func ParseRequestToJson(ctx *gin.Context, request interface{}) (err error) {
	return ctx.ShouldBindBodyWith(&request, binding.JSON)
}

func ValidateEmptyRequest(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}
	return true
}

var ErrorInvalidPassword = errors.New("password did not match")

//get pointer real value
func PointerValue(i interface{}) (value interface{}) {
	var v = reflect.ValueOf(i)
	if !v.IsZero() {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		return v.Interface()
	}
	return nil
}

func MakeSlug(text string) string {

	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}

// rewriting model response
// for i := range teachers {
// 	tempV := &teachers[i]
// 	vGender := lib.PointerValue(tempV.Gender)
// 	// var v = reflect.ValueOf(vGender)
// 	fmt.Println(vGender)
// 	if vGender == uint8(1) {
// 		tempV.Alias = "Laki-Laki"
// 	} else {
// 		tempV.Alias = "Perempuan"
// 	}
// }
