package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

//Setup database instance, it will return gorm database
//dbType: will be database type which supported by GORM
//psn: connection string
func Setup(dbType string, psn string) *gorm.DB {
	db, err := gorm.Open(dbType, psn)
	if err != nil {
		log.Errorf("Oops! Something went wrong, %v", err)
	}
	return db
}
