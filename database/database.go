package database

import (
	"log"
	"os"

	"github.com/Elaborate-backend/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // blank import for gorm
)

var DB *gorm.DB

func init() {
	DB = gormConnect()
	log.Println("[INFO] database connected")

	if !DB.HasTable(&entity.User{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.User{})
		log.Println("[INFO] User table created")
	}

	if !DB.HasTable(&entity.Directory{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Directory{})
		log.Println("[INFO] Directory table created")
	}

	if !DB.HasTable(&entity.Branch{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Branch{})
		log.Println("[INFO] Branch table created")
	}

	if !DB.HasTable(&entity.Commit{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Commit{})
		log.Println("[INFO] Commit table created")
	}

	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.Directory{})
	DB.AutoMigrate(&entity.Branch{})
	DB.AutoMigrate(&entity.Commit{})
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "progate-mafia"
	PASS := "ninjawanko"
	// TODO: production環境も加味
	HOST := map[bool]string{false: "mysql", true: "127.0.0.1"}[os.Getenv("GO_ENV") == "test"]
	PROTOCOL := "tcp(" + HOST + ":3306)"
	DBNAME := "elaborate"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	log.Println("[INFO] should be connected just once.")

	return db
}
