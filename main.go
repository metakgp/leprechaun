package main

import (
	"log"
	"net/http"
    "github.com/joho/godotenv"
    "os"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Print("Error loading .env file")
    }

	router := NewRouter()

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
