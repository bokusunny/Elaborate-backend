package main

import (
	"io/ioutil"
	"log"
	"net/http"
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

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}

		log.Printf("I got post request, json: " + string(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	db := gormConnect()
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}

	db.AutoMigrate(&User{})

	http.HandleFunc("/", postUserHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
