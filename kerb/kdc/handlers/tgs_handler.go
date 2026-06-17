package handlers

import (
	"encoding/base64"
	"encoding/json"
)

func NewTGSHandler(kdc KDC, clientId64, _auth64, _tgt64 string) *TGSHandler {
	return &TGSHandler{
		kdc:        kdc,
		clientId64: clientId64,
		_auth64:    _auth64,
		_tgt64:     _tgt64,
	}
}

type TGSHandler struct {
	kdc                         KDC
	clientId64, _auth64, _tgt64 string
}

func (t *TGSHandler) Handle() ([]byte, error) {
	clientId, _ := base64.StdEncoding.DecodeString(t.clientId64)
	_auth, _ := base64.StdEncoding.DecodeString(t._auth64)
	_tgt, _ := base64.StdEncoding.DecodeString(t._tgt64)

	if _sk2, _tgs, err := t.kdc.GetTGS(clientId, _auth, _tgt); err != nil {
		return nil, err
	} else {
		return json.Marshal(KDCResponse{
			_sk2, _tgs,
		})
	}
}
