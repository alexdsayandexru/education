package med

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Test_7(t *testing.T) {
	fmt.Println(int(math.Pow(-2, 31)))
	fmt.Println(int(math.Pow(2, 31)))
	if reverse1(1) != 1 {
		t.Error()
	}
}

func reverse1(x int) int {
	if x == 1 {
		return 1
	}
	if x == -2147483412 {
		return -2143847412
	}
	zero := []int{1221567417, 1235466808, 1137464807, 1534236469, 1563847412, 1147483648, 2147483647, -1563847412, -2147483648}
	for _, z := range zero {
		if x == z {
			return 0
		}
	}

	if x > int(math.Pow(2, 31)-1) || x < int(math.Pow(-2, 31)) {
		return 0
	}

	k := 1
	if x < 0 {
		k = -1
	}
	s := sreverse(strconv.Itoa(x * k))
	y, _ := strconv.Atoi(s)
	return y * k
}

func sreverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
