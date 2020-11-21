package main

import (
	"fmt"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is Sparta")
}

func HandleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}
