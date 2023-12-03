package repo

import (
	"crypto/rand"
	"database/sql"
	"log"
	"time"

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
	Salt           string
}

func Init() *Repo {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "172.25.0.2:3306",
		DBName:               "login",
		AllowNativePasswords: true,
	}

	var db *sql.DB
	var err error
	for {
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Printf("failed connecting to the database: %s", err.Error())
			time.Sleep(time.Second * 3)
			continue
		}
		break
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Printf("failed to ping the database: %s", pingErr.Error())
	}

	log.Print("Database connection established")

	return &Repo{db}
}

func (r *Repo) AddUser(ld LoginData) {
	encUser, err := ld.encryptLoginData()
	if err != nil {
		log.Printf("failed to encrypt new user data: %s", err.Error())
		return
	}
	_, err = r.db.Exec("INSERT INTO user (username, password, salt) VALUES (?, ?, ?)", encUser.Username, encUser.HashedPassword, encUser.Salt)
	if err != nil {
		log.Printf("failed to add new user: %s", err.Error())
	}
	log.Printf("user: %s added", ld.Username)
}

	err = bcrypt.CompareHashAndPassword([]byte(us.HashedPassword), pswd)
	if err != nil {
		log.Printf("failed to compare passwords: %s", err.Error())
		return err
	}
	log.Print("passwords matched")
	return nil
}

func (r *Repo) fetchUser(ld LoginData) (User, error) {
	row := r.db.QueryRow("SELECT salt, password FROM user WHERE username = ?;", ld.Username)

	us := User{}
	err := row.Scan(&us.Salt, &us.HashedPassword)
	if err != nil {
		log.Printf("failed to fetch user %s: %s", ld.Username, err.Error())
		return us, err
	}
	return us, nil
}

// LoginData to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ld LoginData) encryptLoginData() (User, error) {
	log.Print("encrypting user data")
	var us User
	us.Username = ld.Username
	us.Salt = generateSalt()
	hashedPassword, err := hashPassword(append(us.Salt, []byte(ld.Password)...))
	if err != nil {
		return User{}, err
	}
	us.HashedPassword = hashedPassword
	log.Print(" user data encrypted")
	return us, nil
}

func generateSalt() []byte {
	salt := make([]byte, 4)
	rand.Read(salt)
	return salt
}

func hashPassword(pswd []byte) (string, error) {
	// Generate the bcrypt hash with a cost factor of 10
	log.Print("hashing password")
	hash, err := bcrypt.GenerateFromPassword(pswd, 10)
	if err != nil {
		log.Print("failed hashing password")
		return "", err
	}
	return string(hash), err

}
