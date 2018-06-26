package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	yourname := os.Getenv("YOURNAME")
	fmt.Fprintf(w, "Hello "+yourname+"!!! - "+time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", root)               // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
