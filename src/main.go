package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"loginpage/srv"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// Struct of data to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func encryptLoginData(ld LoginData) srv.User {
	var us srv.User
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

func main() {
	fmt.Println("Listening at port :8080")

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
