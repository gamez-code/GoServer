package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is Sparta")
}

func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}

func PostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
