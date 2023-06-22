package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// Struct of data to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username       string
	HashedPassword string
	Salt           []byte
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
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
	addUser(loginData.Username, nil, loginData.Password)

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

func handleSignup(w http.ResponseWriter, r *http.Request) {
	// extract data
	decoder := json.NewDecoder(r.Body)
	var loginData LoginData

	w.Header().Set("Content-Type", "application/json")

	err := decoder.Decode(&loginData)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// encrypt password
	encryptedData := encryptLoginData(loginData)

	// pass data to datanase
	addUser(encryptedData.Username, encryptedData.Salt, encryptedData.HashedPassword)

	// response
	response := map[string]interface{}{
		"message": "SignUp successful",
	}

	json.NewEncoder(w).Encode(response)
}

func encryptLoginData(ld LoginData) User {
	var us User
	us.Username = ld.Username
	us.Salt = generateSalt()
	us.HashedPassword = hashPassword(ld.Password, us.Salt)
	return us
}

func generateSalt() []byte {
	salt := make([]byte, 4) // Generate a 16-byte (128-bit) salt

	rand.Read(salt)

	return salt
}

func hashPassword(pswd string, salt []byte) string {
	passwordWithSalt := append([]byte(pswd), salt...)

	// Generate the bcrypt hash with a cost factor of 10
	hash, _ := bcrypt.GenerateFromPassword(passwordWithSalt, 10)
	return string(hash)

}

func addUser(username string, salt []byte, password string) {
	_, err := db.Exec("INSERT INTO users (username, salt, password) VALUES (?, ?, ?)", username, salt, password)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Listening at port :8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/signup", handleSignup)

	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/", fs)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
