package database

import (
	"log"

	"github.com/Elaborate-backend/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // blank import for gorm
)

var DB *gorm.DB

func init() {
	DB = gormConnect()
	log.Printf("database connected\n")

	if !DB.HasTable(&entity.User{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.User{})
	}

	DB.AutoMigrate(&entity.User{})
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "progate-mafia"
	PASS := "ninjawanko"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "elaborate"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	log.Printf("should be connected just once.\n")

	return db
}
