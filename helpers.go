package main

import (
    "os"
    "html/template"
    "bytes"
    "time"
    "fmt"
    "crypto/sha256"
    "math/rand"
    "strings"
)

// TODO
func validRoll(roll string) bool {
    return true
}

// TODO
func validEmail(email string) bool {
    return true
}

// ENHANCE: Have different tokens for different applications
func PublicApiAuthenticate(authorization string) bool {
    return authorization == os.Getenv("AUTH_TOKEN")
}

type AuthTemplateContext struct {
    Verifier string
    LinkSuffix string
    BaseLink string
}

func buildAuthPage(verifier string, link_suffix string) string {
    res := AuthTemplateContext{
                    verifier,
                    link_suffix,
                    os.Getenv("BASE_LINK"),
                }
    new_temp, _ := template.ParseFiles(PATH_BEGIN_AUTH_PAGE)
    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}

type AuthUnsuccessfulTemplateContext struct {
    RollExists bool
    EmailExists bool
}

func buildAuthUnsuccessful(rollExists int, emailExists int) string {
    res := AuthUnsuccessfulTemplateContext{
                    rollExists > 0,
                    emailExists > 0,
                }

    new_temp, _ := template.ParseFiles(PATH_BEGIN_AUTH_UNSUCCESSFUL_PAGE)

    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}

type Step1Complete struct {
    Email string
    CompletedTime string
    Completed bool
}

func buildStep1CompletePage(email string, completedTime time.Time, completed bool) string {
    res := Step1Complete{
                    email,
                    completedTime.Format(time.RFC822),
                    completed,
                }
    new_temp, _ := template.ParseFiles(PATH_STEP1_COMPLETE_PAGE)
    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}

type Step2Complete struct {
    Roll string
    Email string
}

func buildStep2CompletePage(roll string, email string) string {
    res := Step2Complete{
                    roll,
                    email,
                }
    new_temp, _ := template.ParseFiles(PATH_STEP2_COMPLETE_PAGE)
    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}

func buildResetPage(email string) string {
    res := struct{
            Email string
        }{
            email,
        }

    new_temp, _ := template.ParseFiles(PATH_BEGIN_RESET_PAGE)
    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}

func buildResetCompletePage(roll string, email string) string {
    res := struct{
        Roll string
        Email string
    }{
        roll, email,
    }

    new_temp, _ := template.ParseFiles(PATH_RESET_COMPLETE_PAGE)
    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}

func getSha256Sum(base string) string {
    h := sha256.New()
    h.Write([]byte(base))
    return fmt.Sprintf("%x", h.Sum(nil))
}

func getSha256SumRandom(base string) string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    newBase := fmt.Sprintf("%s %v %v", base, time.Now().UnixNano(), r.Uint64())

    h := sha256.New()
    h.Write([]byte(newBase))
    return fmt.Sprintf("%x", h.Sum(nil))
}

// ENHANCE: Improve the redaction scheme
func redactEmail(email string) string {
    comps := strings.Split(email, "@")
    username := comps[0]
    domain := comps[1]

    newUsername := username[:]
    if (len(username) > 4) {
        newUsername = strings.Replace(username, username[2:len(username)-2], strings.Repeat("x", len(username)-4), 1)
    }

    newDomain := domain[:]
    if (len(domain) > 4) {
        newDomain = strings.Replace(domain, domain[2:len(domain)-2], strings.Repeat("x", len(domain)-4), 1)
    }
    return fmt.Sprintf("%s@%s", newUsername, newDomain)
}
