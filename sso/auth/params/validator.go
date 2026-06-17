package params

import (
	"errors"
	"fmt"
	"strings"
)

func Valid() *Validator {
	return &Validator{
		ok: true,
	}
}

type Validator struct {
	errs string
	ok   bool
}

func (t *Validator) Required(name string, value string) *Validator {
	if len(strings.Trim(value, " ")) == 0 {
		if t.ok {
			t.errs = fmt.Sprintf("Не найден обязательный параметр в запросе:[%s]", name)
		} else {
			t.errs = t.errs + fmt.Sprintf(",[%s]", name)
		}
		t.ok = false
	}
	return t
}

func (t *Validator) RequiredIf(name string, value string, predicate func() bool) *Validator {
	if len(strings.Trim(value, " ")) == 0 && predicate() {
		if t.ok {
			t.errs = fmt.Sprintf("Не найден обязательный параметр в запросе:[%s]", name)
		} else {
			t.errs = t.errs + fmt.Sprintf(",[%s]", name)
		}
		t.ok = false
	}
	return t
}

func (t *Validator) Result() error {
	if !t.ok {
		return errors.New(t.errs)
	}
	return nil
}
