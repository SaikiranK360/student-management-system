package db

import (
	"github.com/SaikiranK360/grpc-go-practice-2/shared/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MYSQLDB *gorm.DB

func InitiateMySQLDBConnection() {
	var err error
	dsn := config.Config.DB.Username + ":" + config.Config.DB.PassWord + "@tcp" + config.Config.DB.StudentDBUrl + "?charset=" + config.Config.DB.Charset + "&parseTime=True&loc=Local"
	MYSQLDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
