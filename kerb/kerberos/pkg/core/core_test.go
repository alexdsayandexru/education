package core

import (
	"testing"
)

func TestAuthenticator(t *testing.T) {
	crypto := NewEAS()
	key32 := GetHash32("hello")
	if _auth, err := NewAuthenticator(key32, crypto); err != nil {
		t.Error(err)
	} else {
		var auth Authenticator
		if err = DecryptUnmarshal(_auth, key32, crypto, &auth); err != nil {
			t.Error(err)
		}
	}
}
