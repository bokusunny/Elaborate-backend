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

	DB.AutoMigrate(&entity.Directory{}, &entity.Branch{}, &entity.Commit{})
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "elaborate"
	PASS := "ninjawanko0714"
	var HOST string
	switch env := os.Getenv("GO_ENV"); env {
	case "prod":
		HOST = "elaborate-db-instance.cqzfkiclzjux.ap-northeast-1.rds.amazonaws.com"
	case "dev":
		HOST = "mysql"
	case "test":
		HOST = "127.0.0.1"
	}
	PROTOCOL := "tcp(" + HOST + ":3306)"
	DBNAME := "elaborate"
	OPTIONS := "parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTIONS
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	log.Println("[INFO] should be connected just once.")

	return db
}
