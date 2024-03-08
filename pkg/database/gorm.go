package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() (db *gorm.DB) {

	var (
		host     = viper.Get("db.host")
		user     = viper.Get("db.user")
		password = viper.Get("db.password")
		dbname   = viper.Get("db.dbname")
		port     = viper.Get("db.port")
	)

	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok`,
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("can't connect to database: %s", err))
	}

	return db

}
