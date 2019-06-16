package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User has user data sent by firebase cloud functions
type User struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	return db
}

func newUser(name string, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func main() {
	db := gormConnect()
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}

	db.AutoMigrate(&User{})

	createUser := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}

		log.Printf("I got post request, json: " + string(body))

		var user User
		if err := json.Unmarshal(body, &user); err != nil {
			log.Fatal(err)
		}

		newUser := newUser(user.Name, user.Email)
		db.Create(&newUser)
	}

	http.HandleFunc("/", createUser)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
