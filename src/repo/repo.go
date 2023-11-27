package repo

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Repo struct {
	db *sql.DB
}

// User structure
type User struct {
	Username       string
	HashedPassword string
	Salt           []byte
}

func Init() *Repo {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "login",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("failed connecting to the database: %s", err.Error())
		fmt.Printf("failed connecting to the database: %+v", cfg)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("failed to ping the database: %s", pingErr.Error())
	}

	return &Repo{db}
}

func (r *Repo) AddUser(ld LoginData) {
	encUser := ld.encryptLoginData()
	_, err := r.db.Exec("INSERT INTO users (username, salt, password) VALUES (?, ?, ?)", encUser.Username, encUser.Salt, encUser.HashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Repo) GetUser(ld LoginData) error {
	row := r.db.QueryRow("SELECT salt, password FROM users WHERE username = ?;", ld.Username)

	var us User
	err := row.Scan(&us.Salt, &us.HashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	pswd := append(us.Salt, []byte(ld.Password)...)
	return bcrypt.CompareHashAndPassword([]byte(us.HashedPassword), pswd)
}

// LoginData to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ld LoginData) encryptLoginData() User {
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
	passwordWithSalt := append(salt, []byte(pswd)...)

	// Generate the bcrypt hash with a cost factor of 10
	hash, _ := bcrypt.GenerateFromPassword(passwordWithSalt, 10)
	return string(hash)

}
