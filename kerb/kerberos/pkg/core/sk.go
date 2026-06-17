package core

import (
	"time"
)

func NewSK(sk []byte, timestamp time.Time, lifetime time.Time) *SK {
	return &SK{
		sk,
		timestamp,
		lifetime,
	}
}

type SK struct {
	SK        []byte
	Timestamp time.Time
	Lifetime  time.Time
}
