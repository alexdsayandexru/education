package main

import "kerberos/pkg/core"

func NewKDR() KDR {
	return &KDRTest{}
}

type KDRTest struct {
}

func (t *KDRTest) GetPasswordHash32(clientId string) []byte {
	return core.GetHash32(clientId)
}
