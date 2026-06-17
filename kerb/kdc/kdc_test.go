package main

import (
	"fmt"
	"kerberos/pkg/core"
	"testing"
)

func TestKDC(t *testing.T) {
	crypto := core.NewEAS()
	kdcHash32 := core.GetHash32("kdc_password")
	kdc := NewKDC(NewKDR(), kdcHash32, crypto)

	clientHash32 := core.GetHash32("client_password")
	if _auth, err := core.NewAuthenticator(clientHash32, crypto); err != nil {
		t.Error(err)
	} else {
		if _sk1, _tgt, err := kdc.GetTGT([]byte("client_password"), _auth); err != nil {
			t.Error(err)
		} else {
			var sk core.SK
			if err := core.DecryptUnmarshal(_sk1, clientHash32, crypto, &sk); err != nil {
				t.Error(err)
			} else {
				_auth, _ := core.NewAuthenticator(sk.SK, crypto)
				if _sk2, _tgs, err := kdc.GetTGS([]byte("server_password"), _auth, _tgt); err != nil {
					t.Error(err)
				} else {
					serverHash32 := core.GetHash32("server_password")
					var tgs core.TGS
					if err := core.DecryptUnmarshal(_tgs, serverHash32, crypto, &tgs); err != nil {
						t.Error(err)
					} else {
						fmt.Println(tgs, len(_sk2))
					}
				}
			}
		}
	}
}
