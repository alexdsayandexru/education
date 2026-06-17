package handlers

type KDC interface {
	GetTGT(clientId []byte, _auth []byte) (_sk1 []byte, _tgt []byte, err error)
	GetTGS(clientId []byte, _auth []byte, _tgt []byte) (_sk2 []byte, _tgs []byte, err error)
}

type KDCResponse struct {
	SK     []byte
	Ticket []byte
}
