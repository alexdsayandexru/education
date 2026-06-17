package easy

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_190(t *testing.T) {
	/*if reverseBits(43261596) != 964176192 {
		t.Error()
	}*/
	if reverseBits(2) != 1073741824 {
		t.Error()
	}
}

func reverseBits(n int) int {
	buffer := []byte(fmt.Sprintf("%32b", n))
	fmt.Println(string(buffer))
	i := 0
	j := len(buffer) - 1
	for {
		if i >= j {
			break
		}
		if buffer[i] == 32 {
			buffer[i] = 48
		}
		if buffer[j] == 32 {
			buffer[j] = 48
		}

		buffer[i], buffer[j] = buffer[j], buffer[i]
		i++
		j--
	}
	fmt.Println(string(buffer))
	if val, err := strconv.ParseInt(string(buffer), 2, 64); err == nil {
		return int(val)
	}
	return 0
}
