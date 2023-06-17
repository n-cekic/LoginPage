package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

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
	fmt.Printf("Received login data: %+v\n", loginData)

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

func main() {
	fmt.Println("Listening at port :8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleRequest)

	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
