package srv

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type User struct {
	Username       string
	HashedPassword string
	Salt           []byte
}

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/signup", handleSignup)

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

func addUser(username string, salt []byte, password string) {
	_, err := db.Exec("INSERT INTO users (username, salt, password) VALUES (?, ?, ?)", username, salt, password)
	if err != nil {
		log.Fatal(err)
	}
}
