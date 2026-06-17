package client

func NewSecurityContext(clientId string, clientHash []byte, serverId string) *SecurityContext {
	return &SecurityContext{
		clientId:   []byte(clientId),
		clientHash: clientHash,
		serverId:   []byte(serverId),
	}
}

type SecurityContext struct {
	clientId   []byte
	clientHash []byte
	serverId   []byte
	_sk1       []byte
	sk1        []byte
	tgt        []byte
	_sk2       []byte
	sk2        []byte
	tgs        []byte
}

func (t *SecurityContext) ResetTGS() {
	t._sk2 = nil
	t.sk2 = nil
	t.tgs = nil
}

func (t *SecurityContext) ResetTGT() {
	t._sk1 = nil
	t.sk1 = nil
	t.tgt = nil
	t._sk2 = nil
	t.sk2 = nil
	t.tgs = nil
}
