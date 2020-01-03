# go-slack-utils

[![GoDoc](https://godoc.org/github.com/chiefy/go-slack-utils?status.svg)](https://godoc.org/github.com/chiefy/go-slack-utils)

## What this is?

This is a general purpose utility library for using Slack's [Block kit UI](https://api.slack.com/reference/block-kit/blocks), with Go structs corresponding to the blocks used to create UI elements. Also included is middleware for validing Slack requests using HMAC-256 and the Slack secret signing key.


## What this is not?

A Slack API wrapper. There's plenty of those out there.


## Installation

```
go get -u github.com/chiefy/go-slack-utils
```

## Usage

### Middleware

```
func main() {
 	
    r := mux.NewRouter()
	r.HandleFunc("/command", MySlashCommandHandler).Methods(http.MethodPost)
	r.Use(middleware.ValidateTimestamp)

    // It's up to you on how you configure injection of the slack signing secret
    signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
    // Generate the validation middleware by injecting the secret
    validateReq := middleware.ValidateSlackRequest(signingSecret)
	r.Use(validateReq)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
```