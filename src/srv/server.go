package srv

import (
	"encoding/json"
	"fmt"
	"log"
	"loginpage/repo"
	"net/http"
)

// User structure
type User struct {
	Username       string
	HashedPassword string
	Salt           []byte
}

// LoginData to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Server struct {
	repo *repo.Repo
}

func Init() {
	r := repo.Init()
	server := Server{
		repo: r,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", server.handleLogin)
	mux.HandleFunc("/signup", server.handleSignup)

	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/", fs)

	go func() {
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Listening at port :8080")
	}()
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
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
	msg := "Login successful"
	err = s.login(loginData)
	if err != nil {
		fmt.Printf("failed to login: %s", err.Error())
		msg = "Login failed"
	}

	response := map[string]interface{}{
		"message": msg,
	}

	log.Printf("response: %+v", response)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleSignup(w http.ResponseWriter, r *http.Request) {
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

	// pass data to datanase
	s.addUser(loginData)

	// response
	response := map[string]interface{}{
		"message": "SignUp successful",
	}

	log.Printf("response: %+v", response)
	json.NewEncoder(w).Encode(response)
}

func (s Server) addUser(ld LoginData) {
	s.repo.AddUser(repo.LoginData(ld))
}

func (s Server) login(ld LoginData) error {
	return s.repo.GetUser(repo.LoginData(ld))
}
