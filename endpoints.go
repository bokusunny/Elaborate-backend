package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/Elaborate-backend/api"
	"github.com/Elaborate-backend/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile("./service_account_key.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("[ERROR] %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("[ERROR] %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			log.Printf("[ERROR] fail to verify ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		} else {
			// handler内でuidを取得できるように
			r.Header.Set("sub", token.UID)
			log.Printf("[INFO] uid in Verified token: %v\n", token.UID)
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	db := database.DB
	defer db.Close()

	// TODO: originは環境によって場合分け
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

	r := mux.NewRouter()
	r.HandleFunc("/", authMiddleware(api.CreateUserHandler)) // TODO: "/users"に変更 && POSTリクエストに限定
	r.HandleFunc("/directories", authMiddleware(api.CreateDirectoryHandler))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
