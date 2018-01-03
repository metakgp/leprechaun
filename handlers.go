package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
    // "strconv"
	// "github.com/gorilla/mux"
	// "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    b, err := ioutil.ReadFile("index.html")
    if err != nil {
        fmt.Fprintf(w, "Could not read HTML file from disk. Error: %v", err)
    } else {
        fmt.Fprintf(w, "%s", b)
    }
}

func BeginAuth(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    roll := r.PostForm.Get("roll")
    email := r.PostForm.Get("email")

    if !validRoll(roll) {
        fmt.Fprintf(w, "Roll number not valid. Please try again")
        return
    }

    if !validEmail(email) {
        fmt.Fprintf(w, "Email not valid. Please try again")
        return
    }

    fmt.Fprint(w, "You have submitted valid roll and email")
}
