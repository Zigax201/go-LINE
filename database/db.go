package database

import (
	"fmt"
	"go-line/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL(dbConf *config.DatabaseConfiguration) *gorm.DB {
	DBUri := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		dbConf.DBUsername,
		dbConf.DBPassword,
		dbConf.DBIPAddr,
		dbConf.DBName)

	instance, err := gorm.Open(dbConf.DBDialect, DBUri)
	if err != nil {
		log.Panic(fmt.Sprintf("DB Not Connected,\n%s", err))
	}

	return instance
}
