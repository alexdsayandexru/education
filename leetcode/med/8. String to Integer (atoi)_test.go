package med

import (
	"math"
	"strings"
	"testing"
)

func Test_8(t *testing.T) {
	//fmt.Println(int(math.Pow(2, 31)), -2147483648)

	if myAtoi("      -20000000000000000000000000x") != -2147483648 {
		t.Error()
	}

}

func myAtoi(s string) int {
	/*if s == "20000000000000000000" {
		return int(math.Pow(2, 31)) - 1
	}
	if s == "      -20000000000000000000000000x" {
		return -int(math.Pow(2, 31))
	}*/
	sign := 0
	var buffer []byte
	for _, b := range []byte(strings.Trim(s, " ")) {
		if b == '-' && len(buffer) == 0 && sign == 0 {
			sign = -1
		} else if b == '+' && len(buffer) == 0 && sign == 0 {
			sign = 1
		} else if b >= '0' && b <= '9' {
			if sign == 0 {
				sign = 1
			}
			buffer = append(buffer, b)
		} else {
			break
		}
	}
	result := 0
	n := len(buffer) - 1
	for i, b := range buffer {
		result += int(b-'0') * int(math.Pow(10, float64(n-i)))
	}

	if result < 0 {
		result = int(math.Pow(2, 31))
	}

	result *= sign

	Max := int(math.Pow(2, 31))
	if result >= Max {
		return Max - 1
	}
	Min := -int(math.Pow(2, 31))
	if result <= Min {
		return Min
	}
	return result
}
