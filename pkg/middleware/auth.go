package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	slackSigHeader          = "X-Slack-Signature"
	slackSigHeaderTimestamp = "X-Slack-Request-Timestamp"
	// RequestTTL asserts incoming requests must have a timestamp within this duration from server's time
	RequestTTL = "1m"
)

// ValidateTimestamp asserts that the incoming request's timestamp is within a reasonable timespan from current time
func ValidateTimestamp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !validateTimestamp(timestampFromRequest(r)) {
			log.Println("invalid timestamp")
			http.Error(w, "Invalid Timestamp", http.StatusBadRequest)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// timestampFromRequest gets the unix epoch timestamp on the http request
func timestampFromRequest(r *http.Request) int64 {
	ts, err := strconv.ParseInt(r.Header.Get(slackSigHeaderTimestamp), 10, 64)
	if err != nil {
		log.Println("could not get timestamp", err)
		ts = 0
	}
	return ts
}

// validateTimestamp asserts that the request's timestamp is within a reasonable timeframe compared with server's current time
func validateTimestamp(ts int64) bool {
	abs := func(a time.Duration) time.Duration {
		if a >= 0 {
			return a
		}
		return -a
	}
	m, _ := time.ParseDuration(RequestTTL)
	d := abs(time.Since(time.Unix(ts, 0)))
	return d < m
}

// ValidateSlackRequest validates a request's signature is signed by the provided slack secret token
func ValidateSlackRequest(signingSecretToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			bodyData, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("bad request body %s", err)
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			incomingSig := []byte(fmt.Sprintf("v0:%d:%s", timestampFromRequest(r), string(bodyData)))
			slackSig, _ := hex.DecodeString(strings.TrimPrefix(r.Header.Get(slackSigHeader), "v0="))
			if !validMAC(incomingSig, slackSig, []byte(signingSecretToken)) {
				log.Println("HMAC error - signatures did not match")
				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				r.Body.Close()
				r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))
				next.ServeHTTP(w, r)
			}
		}
		return http.HandlerFunc(f)
	}
}

// validMAC reports whether messageMAC is a valid HMAC tag for message.
func validMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
