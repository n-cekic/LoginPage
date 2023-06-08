package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct of data to be recieved
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginData LoginData

	err := decoder.Decode(&loginData)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Process loginData
	fmt.Printf("Received login data: %+v\n", loginData)

	// Send response
	response := map[string]interface{}{
		"message": "Login successful",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Listening at port :8080")
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
