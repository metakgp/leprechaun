package main

import (
    "os"
    "html/template"
    "bytes"
    "time"
)

// TODO
func validRoll(roll string) bool {
    return true
}

// TODO
func validEmail(email string) bool {
    return true
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
    new_temp, _ := template.ParseFiles("begin_auth.tmpl.html")
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
    new_temp, _ := template.ParseFiles("step1_complete.tmpl.html")
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
    new_temp, _ := template.ParseFiles("step2_complete.tmpl.html")
    var templated_res bytes.Buffer
    new_temp.Execute(&templated_res, res)
    return templated_res.String()
}
