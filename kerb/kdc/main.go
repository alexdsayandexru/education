package main

import (
	"errors"
	"kdc/handlers"
	"kerberos/pkg/core"
	"log"
	"net/http"
)

const (
	ClientId      = "client_id"
	Authenticator = "authenticator"
	Password      = "helloworld"
	Tgt           = "tgt"
)

func GetKDC() *KDC {
	return NewKDC(NewKDR(), core.GetHash32(Password), core.NewEAS())
}

func main() {
	http.HandleFunc("/kdc/tgt", tgt)
	http.HandleFunc("/kdc/tgs", tgs)

	log.Fatal(http.ListenAndServe(":8888", nil))
}

type Handler interface {
	Handle() ([]byte, error)
}

func tgt(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	handle(handlers.NewTGTHandler(GetKDC(), r.Form.Get(ClientId), r.Form.Get(Authenticator)), w)
}

func tgs(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	handle(handlers.NewTGSHandler(GetKDC(), r.Form.Get(ClientId), r.Form.Get(Authenticator), r.Form.Get(Tgt)), w)
}

func handle(h Handler, w http.ResponseWriter) {
	if resp, err := h.Handle(); err != nil {
		var e *core.KerbError
		if errors.As(err, &e) {
			http.Error(w, err.Error(), e.Code())
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		_, _ = w.Write(resp)
	}
}
