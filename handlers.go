package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
    "os"
    "log"
    "net/url"
    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2/bson"
    "strings"
    "time"
)

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    b, err := ioutil.ReadFile(PATH_INDEX_PAGE)
    if err != nil {
        fmt.Fprintln(w, "Could not read HTML file from disk. Error: ", err)
        d, _ := os.Getwd()
        fmt.Fprintf(w, "Currently in %s, searching for %s", d, PATH_INDEX_PAGE)
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

    p := GetPerson(roll, email)

    c := GlobalDBSession.DB(os.Getenv("DB_NAME")).C("people")
    err := c.Insert(&p)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "%s", buildAuthPage(p.Verifier, p.LinkSuffix))
}

func getSingleSecQues(roll string) string {
    v := url.Values{}
    v.Set("user_id", roll)
    resp, _ := http.PostForm(ERP_SECRET_QUES_URL, v)
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}

func getSecurityQuestions(roll string) []string {
    allSecQues := []string{}

    // Perform upto 30 tries to get the 3 unique secret questions from ERP
    for i := 1; i < 30; i++ {
        secQues := getSingleSecQues(roll)
        log.Printf("Run %d, Got %s", i, secQues)
        alreadyFound := false
        for _, q := range allSecQues {
            if q == secQues {
                alreadyFound = true
                break;
            }
        }

        if !alreadyFound {
            allSecQues = append(allSecQues[:], secQues)
        }

        if len(allSecQues) >= 3 {
            break;
        }
    }
    return allSecQues;
}

func VerifyStep1(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    linkSuf := vars["token"]

    c := GlobalDBSession.DB(os.Getenv("DB_NAME")).C("people")

    var result Person
    err := c.Find(bson.M{"linksuffix": linkSuf}).One(&result)

    if err != nil {
        fmt.Fprint(w, "That verifier token isn't there in our DB!")
        return
    }

    verified := false

    if !result.Step1Complete {

        secQues := getSecurityQuestions(result.Roll)

        for _, ques := range secQues {
            if strings.Contains(ques, result.Verifier) {
                verified = true
                break
            }
        }
    }

    if verified || result.Step1Complete {
        fmt.Fprint(w, buildStep1CompletePage(result.Email, result.Step1CompletedAt, result.Step1Complete))
        SendVerificationEmail(result.Email, result.EmailToken)
        if !result.Step1Complete {
            c.Update(bson.M{"linksuffix": linkSuf}, bson.M{ "$set": bson.M{"step1complete": true, "step1completedat": time.Now()} })
        }
    } else {
        fmt.Fprint(w, "Not verified! Go into your ERP and ensure that you have put your verifier token in one of the secret questions!")
    }
}

func VerifyStep2(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    emailTok := vars["token"]
    c := GlobalDBSession.DB(os.Getenv("DB_NAME")).C("people")

    var result Person
    err := c.Find(bson.M{"emailtoken": emailTok}).One(&result)

    if err != nil {
        fmt.Fprint(w, "That email token isn't there in the DB!")
        return
    }

    fmt.Fprint(w, buildStep2CompletePage(result.Roll, result.Email))
    c.Update(bson.M{"emailtoken": emailTok}, bson.M{ "$set": bson.M{"step2complete": true} })
}
