package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
    // "strconv"
	// "github.com/gorilla/mux"
	// "encoding/json"
    // "gopkg.in/mgo.v2
    // "gopkg.in/mgo.v2/bson"
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

    for _, f := range fields {
        if !f.validator(r.PostForm.Get(f.key)) {
            fmt.Fprintf(w, "%s field isn't valid!", f.key)
            return
        }
    }

    roll := r.PostForm.Get("roll")
    email := r.PostForm.Get("email")

    t := GetPerson(roll, email)

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "%s", buildAuthPage(t.verifier, t.link_suffix))
}
