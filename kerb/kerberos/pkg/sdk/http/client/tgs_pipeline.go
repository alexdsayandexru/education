package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"kerberos/pkg/core"
	http_ "kerberos/pkg/sdk/http/http"
	"net/http"
	"net/url"
	"strings"
)

type SuccessCallback func(tgs []byte, auth []byte)
type FailedCallback func(err error)

func NewTGSPipeline(tgtUri string, tgsUri string, ctx *SecurityContext, crypto core.Crypto,
	success SuccessCallback, failed FailedCallback) *TGSPipeline {
	return &TGSPipeline{
		tgtUri:  tgtUri,
		tgsUri:  tgsUri,
		ctx:     ctx,
		crypto:  crypto,
		success: success,
		failed:  failed,
	}
}

type TGSPipeline struct {
	tgtUri  string
	tgsUri  string
	ctx     *SecurityContext
	crypto  core.Crypto
	err     error
	success SuccessCallback
	failed  FailedCallback
}

func (t *TGSPipeline) decrypt(_sk, secretKey []byte) (*core.SK, error) {
	var sk core.SK
	b, err := t.crypto.Decrypt(_sk, secretKey)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(b, &sk)

	return &sk, nil
}

func (t *TGSPipeline) TGT() *TGSPipeline {
	if t.err != nil || t.ctx.tgt != nil {
		return t
	}

	_auth, err := core.NewAuthenticator(t.ctx.clientHash, t.crypto)
	if err != nil {
		t.err = errors.New(fmt.Sprintf("Invalid authenticator:[%s]", err.Error()))
	} else {
		data := url.Values{}
		data.Set("client_id", base64.StdEncoding.EncodeToString(t.ctx.clientId))
		data.Set("authenticator", base64.StdEncoding.EncodeToString(_auth))

		r, _ := http.NewRequest(http.MethodPost, t.tgtUri, strings.NewReader(data.Encode()))
		resp, body := http_.HttpPost(r)

		if resp.StatusCode == http.StatusOK {
			resp := KDCResponse{}
			_ = json.Unmarshal(body, &resp)
			t.ctx._sk1 = resp.SK
			t.ctx.tgt = resp.Ticket
			t.ctx.sk1 = nil
		} else {
			t.err = errors.New(string(body))
		}
	}
	return t
}

func (t *TGSPipeline) SK1() *TGSPipeline {
	if t.err != nil || t.ctx.sk1 != nil {
		return t
	}

	if sk1, err := t.decrypt(t.ctx._sk1, t.ctx.clientHash); err == nil {
		t.ctx.sk1 = sk1.SK
	} else {
		t.err = err
	}
	return t
}

func (t *TGSPipeline) TGS() *TGSPipeline {
	if t.err != nil || t.ctx.tgs != nil {
		return t
	}

	_auth, err := core.NewAuthenticator(t.ctx.sk1, t.crypto)
	if err != nil {
		t.err = errors.New(fmt.Sprintf("Invalid authenticator:[%s]", err.Error()))
	} else {
		data := url.Values{}
		data.Set("client_id", base64.StdEncoding.EncodeToString(t.ctx.serverId))
		data.Set("tgt", base64.StdEncoding.EncodeToString(t.ctx.tgt))
		data.Set("authenticator", base64.StdEncoding.EncodeToString(_auth))

		r, _ := http.NewRequest(http.MethodPost, t.tgsUri, strings.NewReader(data.Encode()))
		resp, body := http_.HttpPost(r)

		if resp.StatusCode == http.StatusOK {
			resp := KDCResponse{}
			_ = json.Unmarshal(body, &resp)
			t.ctx._sk2 = resp.SK
			t.ctx.tgs = resp.Ticket
			t.ctx.sk2 = nil
		} else {
			errInfo := string(body)
			t.err = errors.New(errInfo)
		}
	}
	return t
}

func (t *TGSPipeline) SK2() *TGSPipeline {
	if t.err != nil || t.ctx.sk2 != nil {
		return t
	}

	if sk2, err := t.decrypt(t.ctx._sk2, t.ctx.sk1); err == nil {
		t.ctx.sk2 = sk2.SK
	} else {
		t.err = err
	}
	return t
}

func (t *TGSPipeline) Do() {
	if t.err != nil {
		t.failed(t.err)
	} else {
		if _authenticator, err := core.NewAuthenticator(t.ctx.sk2, t.crypto); err != nil {
			t.err = err
			t.failed(t.err)
		} else {
			t.success(t.ctx.tgs, _authenticator)
		}
	}
}
