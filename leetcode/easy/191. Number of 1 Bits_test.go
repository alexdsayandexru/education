package easy

import (
	"fmt"
	"testing"
)

func Test_191(t *testing.T) {
	if hammingWeight(2147483645) != 30 {
		t.Error()
	}
}

func hammingWeight(n int) int {
	count := 0
	buffer := []byte(fmt.Sprintf("%b", n))
	for _, v := range buffer {
		count += int(v) - 48
	}
	return count
}
