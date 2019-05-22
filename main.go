package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func gormConnect() *gorm.DB {
	// TODO: 環境変数に置き換える
	DBMS := "mysql"
	USER := "root"
	PASS := "Gakufu1407"
	DBNAME := "Elaborate"

	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}

	db.AutoMigrate(&User{})
}
