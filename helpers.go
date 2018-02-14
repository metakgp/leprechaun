package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"strings"
	"time"
	"net/http"
	"errors"
)

// A function that will validate the given roll number string
// Returns true if the given roll number is valid
// TODO
func validRoll(roll string) bool {
	return true
}

// A function to validate the given email address.
// Returns true if the given string is valid
// TODO
func validEmail(email string) bool {
	return true
}

// This function validates the user using the Authorization header that they
// have sent alongwith their request to the API
// ENHANCE: Have different tokens for different applications
func authenticateRequest(req *http.Request) error {
	if req.Header.Get("Authorization") == os.Getenv("AUTH_TOKEN") {
		return nil
	} else {
		return errors.New("Unauthorized")
	}
}

// AuthTemplateContext is the context structure that must be sent to generate
// the markup for the HTML page shown to the user when step1 is initiated
type AuthTemplateContext struct {
	Verifier   string
	LinkSuffix string
	BaseLink   string
}

// BuildAuthPage returns the HTML markup for the page shown to the user when the
// first step of authentication has been initiated by the user
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

// Context structure for page shown when step 1 could not be initiated
type AuthUnsuccessfulTemplateContext struct {
	RollExists  bool
	EmailExists bool
}

// This function returns the HTML markup for the page shown to the user when
// step 1 of the authentication process could not be initiated
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

// Context structure for page shown to the user when step 1 is complete
type Step1Complete struct {
	Email         string
	CompletedTime string
	Completed     bool
}

// Function that generates markup for the page shown to the user when step 1 is
// complete
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

// Context structure for page shown to the user when step 2 is complete
type Step2Complete struct {
	Roll  string
	Email string
}

// Function that generates markup for page shown to the user when step 2 is
// complete
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

// Function to generate the markup of the page shown to the user when they have
// entered either a Roll number or email address which they wish to reset. This
// email string is actually redacted using the redactEmail function in this file
func buildResetPage(email string) string {
	res := struct {
		Email string
	}{
		email,
	}

	new_temp, _ := template.ParseFiles(PATH_BEGIN_RESET_PAGE)
	var templated_res bytes.Buffer
	new_temp.Execute(&templated_res, res)
	return templated_res.String()
}

// Function to generate the markup of the page shown to the user when they have
// succesfully completed resetting their roll number and email ID using the
// reset flow
func buildResetCompletePage(roll string, email string) string {
	res := struct {
		Roll  string
		Email string
	}{
		roll, email,
	}

	new_temp, _ := template.ParseFiles(PATH_RESET_COMPLETE_PAGE)
	var templated_res bytes.Buffer
	new_temp.Execute(&templated_res, res)
	return templated_res.String()
}

// Function to return the SHA256 sum of the given string
func getSha256Sum(base string) string {
	h := sha256.New()
	h.Write([]byte(base))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Function to return the SHA256 sum of the Base string + present time in Unix
// Nano int64 + a random 64-bit unsigned int generated using the math/rand
// module
func getSha256SumRandom(base string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	newBase := fmt.Sprintf("%s %v %v", base, time.Now().UnixNano(), r.Uint64())

	h := sha256.New()
	h.Write([]byte(newBase))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Function to redact the given email string. It shows the first 2 and last 2
// characters of the username and the domain and replaces everything else with
// "x".
//
// If the email is too small, then nothing might be redacted.
// ENHANCE: Improve the redaction scheme
func redactEmail(email string) string {
	comps := strings.Split(email, "@")
	username := comps[0]
	domain := comps[1]

	newUsername := username[:]
	if len(username) > 4 {
		newUsername = strings.Replace(username, username[2:len(username)-2], strings.Repeat("x", len(username)-4), 1)
	}

	newDomain := domain[:]
	if len(domain) > 4 {
		newDomain = strings.Replace(domain, domain[2:len(domain)-2], strings.Repeat("x", len(domain)-4), 1)
	}
	return fmt.Sprintf("%s@%s", newUsername, newDomain)
}
