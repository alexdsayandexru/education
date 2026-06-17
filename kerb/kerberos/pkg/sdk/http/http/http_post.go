package http

import (
	"io"
	"net/http"
)

const (
	ExpiredTgs  = 462
	ExpiredAuth = 463
	ExpiredTgt  = 464
)

func HttpPost(r *http.Request) (*http.Response, []byte) {
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	if resp, err := client.Do(r); err != nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, []byte(err.Error())
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return resp, bodyBytes
	}
}
