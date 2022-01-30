package repository

import (
	"web_go/database"
	"web_go/model"
)

func FindAllTeachers(teacher *[]model.Teacher) (err error) {
	if err = database.Db.Find(teacher).Error; err != nil {
		return err
	}
	return nil
}
