package repository

import (
	"web_go/database"
	"web_go/model"
)

func FindAllMembers(member *[]model.Member) (err error) {

	if err = database.Db.Find(member).Error; err != nil {
		return err
	}
	return nil
}

func FindAllMembersWithUser(member *[]model.ResponseMemberUser, model *model.MemberUser) (err error) {
	if err = database.Db.Debug().Select("Id", "FullName").Model(model).Preload("User").Find(member).Error; err != nil {
		return err
	}
	return nil

}
