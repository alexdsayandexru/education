package easy

import (
	"fmt"
	"math"
	"testing"
)

func Test_171(t *testing.T) {
	fmt.Println(titleToNumber("ABCD"))
}

func titleToNumber(columnTitle string) int {
	number := 0
	for i, c := range columnTitle {
		number += int(c-'A'+1) * int(math.Pow(26, float64(len(columnTitle)-1-i)))
	}
	return number
}
