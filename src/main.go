package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Struct of data to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	res, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(res))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginData LoginData

	w.Header().Set("Content-Type", "application/json")

	err := decoder.Decode(&loginData)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Process loginData
	fmt.Printf("Adding %+v to DB\n", loginData.Username)
	addUser(loginData.Username, "", loginData.Password)

	// print the request
	res, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(res))

	response := map[string]interface{}{
		"message": "Login successful",
	}

	json.NewEncoder(w).Encode(response)
}

func addUser(username string, email string, password string) {
	if email == "" {
		email = "TEMP_ADDRESS"
	}
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Listening at port :8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleRequest)

	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/", fs)

	ctx, _ := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "users",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	row := db.QueryRow("SELECT * FROM users WHERE UserID = 1;")

	var id int
	var uname string
	var email string
	var pswd string
	row.Scan(&id, &uname, &email, &pswd)
	fmt.Println(id, uname, email, pswd)

	fmt.Println("Connected to the DB!")
	<-ctx.Done()
}
