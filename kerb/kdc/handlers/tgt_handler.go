package handlers

import (
	"encoding/base64"
	"encoding/json"
)

func NewTGTHandler(kdc KDC, clientId64 string, _auth64 string) *TGTHandler {
	return &TGTHandler{
		kdc:        kdc,
		clientId64: clientId64,
		_auth64:    _auth64,
	}
}

type TGTHandler struct {
	kdc                 KDC
	_auth64, clientId64 string
}

func (t *TGTHandler) Handle() ([]byte, error) {
	_auth, _ := base64.StdEncoding.DecodeString(t._auth64)
	clientId, _ := base64.StdEncoding.DecodeString(t.clientId64)

	if _sk1, _tgt, err := t.kdc.GetTGT(clientId, _auth); err != nil {
		return nil, err
	} else {
		b, err := json.Marshal(KDCResponse{
			_sk1, _tgt,
		})
		return b, err
	}
}
