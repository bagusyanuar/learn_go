package repository

import (
	"web_go/database"
	"web_go/lib"
	"web_go/model"
)

func RegisterMember(user *model.UserMember) (err error) {
	if err = database.Db.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}

func SignInMember(user *model.UserMember, email string, password string, provider string) (u *model.UserMember, err error) {
	//JOIN ON change default LEFT JOIN on Preload
	if err = database.Db.Debug().Preload("Member").Joins("JOIN member ON users.id = member.user_id").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	//if member using app provider do check password
	if provider == "app" && !lib.IsPasswordValid(password, user.Password) {
		return user, lib.ErrorInvalidPassword
	}

	return user, nil
}
