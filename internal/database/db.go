package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"hr-system/internal/config"
)

func InitDB() *gorm.DB {
	dsn := config.Get().Database.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
