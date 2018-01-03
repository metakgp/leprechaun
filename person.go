package main

// import "gopkg.in/mgo.v2/bson"
import (
    "time"
    "math/rand"
    "crypto/sha256"
    "fmt"
)

type Person struct {
    // ID bson.ObjectID `bson:"_id,omitempty"`
    Roll string
    Email string
    Verifier string
    EmailToken string
    LinkSuffix string
}

// ENHANCE: Improve the generation of the random seeds
func GetPerson(roll string, email string) Person {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    base := fmt.Sprintf("%s %s %v", roll, email, time.Now().UnixNano())

    h := sha256.New()

    h.Write([]byte(base))

    h.Write([]byte(fmt.Sprintf("%d", r.Uint64())))
    link_suffix := fmt.Sprintf("%x", h.Sum(nil))

    h.Write([]byte(fmt.Sprintf("%d", r.Uint64())))
    verifier := fmt.Sprintf("%x", h.Sum(nil))

    h.Write([]byte(fmt.Sprintf("%d", r.Uint64())))
    email_tok := fmt.Sprintf("%x", h.Sum(nil))

    return Person{
        roll,
        email,
        verifier[:15],
        email_tok[:15],
        link_suffix[:15],
    }
}
