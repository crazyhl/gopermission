package bootstrap

import (
	"github.com/crazyhl/gopermission/base_struct"
	"github.com/crazyhl/gopermission/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Db(dbConfig base_struct.DbConfig) *gorm.DB {
	dsn := dbConfig.Username + ":" + dbConfig.Password +
		"@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")" +
		"/" + dbConfig.Database + "?charset=" + dbConfig.Charset +
		"&parseTime=True&loc=" + dbConfig.Location
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		//DontSupportRenameIndex:  true,
		//DontSupportRenameColumn: true,
		DefaultStringSize: 191,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//_ = db.AutoMigrate(&models.User{})
	config.GetConfig().GormDb = db
	return db
}
