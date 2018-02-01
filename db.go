package main

import (
	"gopkg.in/mgo.v2"
	"os"
)

// DialDB() returns a mgo Session connected to the Database pointed to by the
// DB_URL environment variable
func DialDB() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic
	// behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}
