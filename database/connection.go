package database

import (
	"strconv"

	"gorm.io/gorm"
)

var Db *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func BuildDB() *DBConfig {
	config := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "",
		DBName:   "learn_go",
	}
	return &config
}

func DBUrl(dbconfig *DBConfig) string {
	return dbconfig.User + ":" + dbconfig.Password + "@tcp(" + dbconfig.Host + ":" + strconv.Itoa(dbconfig.Port) + ")/" + dbconfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
