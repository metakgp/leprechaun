package main

import (
    "os"
    "html/template"
    "bytes"
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
