package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

const (
	applicationJSON = "application/json"
)

// GetTestHandler returns a http.HandlerFunc for testing http middleware
func GetTestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		return
	}
	return http.HandlerFunc(fn)
}

func TestValidateTimestampHandler(t *testing.T) {
	assert := assert.New(t)

	tenMinutesAgo, _ := time.ParseDuration("-10m")
	justOverMinute, _ := time.ParseDuration(RequestTTL)
	justOverMinute = justOverMinute + (time.Second * 10)

	tests := []struct {
		description  string
		url          string
		expectedBody string
		expectedCode int
		timestamp    int64
	}{
		{
			description:  "invalid timestamp header",
			url:          "/",
			expectedBody: "Invalid Timestamp\n",
			expectedCode: 400,
			timestamp:    0,
		},
		{
			description:  "ten minutes ago",
			url:          "/",
			expectedBody: "Invalid Timestamp\n",
			expectedCode: 400,
			timestamp:    time.Now().Add(tenMinutesAgo).Unix(),
		},
		{
			description:  "three minutes in the future",
			url:          "/",
			expectedBody: "Invalid Timestamp\n",
			expectedCode: 400,
			timestamp:    time.Now().Add(justOverMinute).Unix(),
		},
		{
			description:  "valid timestamp",
			url:          "/",
			expectedBody: "",
			expectedCode: 200,
			timestamp:    time.Now().Unix(),
		},
	}

	ts := httptest.NewServer(ValidateTimestamp(GetTestHandler()))
	defer ts.Close()

	for _, tc := range tests {
		j, _ := json.Marshal("")

		req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(j))
		assert.NoError(err)

		req.Header.Set("Content-Type", applicationJSON)
		if tc.timestamp != 0 {
			req.Header.Set(slackSigHeaderTimestamp, strconv.Itoa(int(tc.timestamp)))
		}

		res, err := http.DefaultClient.Do(req)
		assert.NoError(err)

		if res != nil {
			defer res.Body.Close()
		}

		b, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)

		assert.Equal(tc.expectedCode, res.StatusCode, tc.description)
		assert.Equal(tc.expectedBody, string(b), tc.description)
	}
}
