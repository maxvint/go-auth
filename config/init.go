package config

import "github.com/jinzhu/gorm"

var (
	DBConfig AppConfig
	TodoDB   *gorm.DB
)

func init() {
	TodoDB, err = DBConfig.DBAddress.DB()
}
