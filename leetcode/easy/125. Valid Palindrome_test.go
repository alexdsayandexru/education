package easy

import (
	"fmt"
	"testing"
)

func Test_125(t *testing.T) {
	//fmt.Println(string([]byte{58, 59, 60, 61, 62, 63, 64}))
	//s := "a."
	//s := "0P"
	//s := "A man, a plan, a canal: Panama"
	//s := "aA"
	s := "Zeus was deified, saw Suez."

	if isPalindrome(s) == false {
		t.Error()
	}
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	var left uint8
	var right uint8

	for {
		if i >= j {
			break
		} else if len(s) == 2 {
			if !isIgnore(s[i]) && !isIgnore(s[j]) {
				return isEqual(s[i], s[j])
			}
			return true
		}
		left = s[i]
		if isIgnore(left) {
			i++
			continue
		}
		right = s[j]
		if isIgnore(right) {
			j--
			continue
		}
		if !isEqual(left, right) {
			fmt.Println(string([]byte{left, right}), left, right, i, j)
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func isEqual(left, right uint8) bool {
	if left >= 65 && left <= 90 {
		left += 32
	}
	if right >= 65 && right <= 90 {
		right += 32
	}
	return left == right
}

func isIgnore(c uint8) bool {
	return c <= 47 || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || c >= 123
}
