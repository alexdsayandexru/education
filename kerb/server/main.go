package main

import (
	"encoding/base64"
	"errors"
	"kerberos/pkg/core"
	"kerberos/pkg/sdk/http/server"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	Authenticator = "authenticator"
	Tgs           = "tgs"
	Password      = "Password"
)

var client *server.KrbClient

func main() {
	client = server.NewKrbClient(core.GetHash32(os.Getenv(Password)), core.NewEAS(), time.Second)

	http.HandleFunc("/server/echo", echo)

	log.Fatal(http.ListenAndServe(":9999", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	handle(r,
		func() {
			_, _ = w.Write([]byte("OK"))
		}, func(err error, code int) {
			http.Error(w, err.Error(), code)
		})
}

func handle(r *http.Request, success func(), failed func(error, int)) {
	_auth, _ := base64.StdEncoding.DecodeString(r.Header.Get(Authenticator))
	_tgs, _ := base64.StdEncoding.DecodeString(r.Header.Get(Tgs))

	if err := client.Validate(_auth, _tgs); err != nil {
		var e *core.KerbError
		if errors.As(err, &e) {
			failed(err, e.Code())
		} else {
			failed(err, 555)
		}
	} else {
		success()
	}
}
