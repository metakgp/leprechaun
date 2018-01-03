package main

import (
    "gopkg.in/mgo.v2"
    "os"
)

func DialDB() *mgo.Session {
    session, err := mgo.Dial(os.Getenv("DB_URL"))
    if err != nil {
        panic(err)
    }

    // Optional. Switch the session to a monotonic
    // behavior.
    session.SetMode(mgo.Monotonic, true)

    return session
}
