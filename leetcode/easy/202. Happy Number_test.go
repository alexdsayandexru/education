package easy

import (
	"strconv"
	"testing"
)

func Test_202(t *testing.T) {
	if isHappy(2) != true {
		t.Error()
	}
}

func isHappy(n int) bool {
	ok := false
	_isHappy(n, &ok, make(map[int]interface{}))
	return ok
}

func _isHappy(n int, ok *bool, m map[int]interface{}) {
	sum := 0
	for _, v := range []byte(strconv.Itoa(n)) {
		sum += (int(v) - 48) * (int(v) - 48)
	}

	if sum != 1 {
		if m[sum] != nil {
			*ok = false
		} else {
			m[sum] = sum
			_isHappy(sum, ok, m)
		}
	} else {
		*ok = true
	}
}
