package client

import (
	"encoding/base64"
	"errors"
	"kerberos/pkg/core"
	http_ "kerberos/pkg/sdk/http/http"
	"log"
	"net/http"
)

const (
	TGTUri        = "http://localhost:8888/kdc/tgt"
	TGSUri        = "http://localhost:8888/kdc/tgs"
	Authenticator = "authenticator"
	Tgs           = "tgs"
)

func NewKrbClient(ctx *SecurityContext, crypto core.Crypto) *KrbClient {
	return &KrbClient{
		ctx:    ctx,
		crypto: crypto,
	}
}

type KrbClient struct {
	ctx    *SecurityContext
	crypto core.Crypto
}

func (t *KrbClient) HttpPost(r *http.Request) (int, string) {
	resp, body := t.httpPost(r)
	if resp.StatusCode == http_.ExpiredTgs {
		log.Println(resp.StatusCode, "TGS")
		t.ctx.ResetTGS()
		resp, body = t.httpPost(r)
	}
	if resp.StatusCode == http_.ExpiredTgt {
		log.Println(resp.StatusCode, "TGT")
		t.ctx.ResetTGT()
		resp, body = t.httpPost(r)
	}
	return resp.StatusCode, string(body)
}

func (t *KrbClient) httpPost(r *http.Request) (resp *http.Response, body []byte) {
	pipeline := NewTGSPipeline(TGTUri, TGSUri, t.ctx, t.crypto,
		func(tgs []byte, auth []byte) {
			r.Header.Set(Tgs, base64.StdEncoding.EncodeToString(tgs))
			r.Header.Set(Authenticator, base64.StdEncoding.EncodeToString(auth))

			resp, body = http_.HttpPost(r)
		},
		func(err error) {
			var e *core.KerbError
			if errors.As(err, &e) {
				resp, body = &http.Response{StatusCode: e.Code()}, []byte(err.Error())
			} else {
				resp, body = &http.Response{StatusCode: 500}, []byte(err.Error())
			}

			//resp, body = &http.Response{StatusCode: http_.ExpiredTgt}, []byte(err.Error())
			resp, body = &http.Response{StatusCode: 500}, []byte(err.Error())
		})

	pipeline.TGT().SK1().TGS().SK2().Do()
	return
}

type KDCResponse struct {
	SK     []byte
	Ticket []byte
}
