package base_struct

import "gorm.io/gorm"

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
	Location string
}

type Config struct {
	DbConfig DbConfig
	GormDb   *gorm.DB
}
