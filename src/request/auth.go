package request

import "mime/multipart"

type AdminSignRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type signUpRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Provider string `form:"provider"`
}

type MemberSignUpRequest struct {
	signUpRequest
	Phone   string `form:"phone"`
	School  string `form:"school"`
	Address string `form:"address"`
	Avatar  string `form:"avatar"`
}

type MemberSignInRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Provider string `form:"provider"`
}

type CategoryRequest struct {
	Name string                `form:"name"`
	Icon *multipart.FileHeader `form:"icon"`
}
