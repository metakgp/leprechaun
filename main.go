package leprechaun

import (
	"log"
	"net/http"
    "github.com/joho/godotenv"
    "os"
)

var GlobalDBSession = DialDB()

func Main() {

    err := godotenv.Load()
    if err != nil {
        log.Print("Error loading .env file")
    }

	router := NewRouter()
    router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

    log.Printf("Server started on %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))

    defer GlobalDBSession.Close()
}
