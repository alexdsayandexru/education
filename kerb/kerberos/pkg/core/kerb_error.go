package core

import "fmt"

func NewKerbError(err string, code int) *KerbError {
	return &KerbError{
		err:  err,
		code: code,
	}
}

type KerbError struct {
	err  string
	code int
}

func (t *KerbError) Error() string {
	return fmt.Sprintf("%s : %d", t.err, t.code)
}

func (t *KerbError) Code() int {
	return t.code
}
