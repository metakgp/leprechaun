package main

import (
	"os"

	"gopkg.in/mgo.v2"
)

// DialDB() returns a mgo Session connected to the Database pointed to by the
// DB_URL environment variable
func DialDB() *mgo.Session {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI := os.Getenv("ATLAS_MONGODB_URI")
	}
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic
	// behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}
