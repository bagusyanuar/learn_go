package repository

import (
	"web_go/database"
	"web_go/model"
)

func FindAllUsers(user *[]model.User) (err error) {
	if err = database.Db.Joins("JOIN member on users.id = member.user_id").Preload("Member").Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *model.User) (err error) {
	if err = database.Db.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserMember(user *model.UserMember) (err error){
	if err = database.Db.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}
