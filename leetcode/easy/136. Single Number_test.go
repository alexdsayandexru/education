package easy

import (
	"fmt"
	"testing"
)

func Test_136(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		v := m[n] + 1
		if v == 2 {
			delete(m, n)
		} else {
			m[n] = v
		}
	}
	for n, s := range m {
		if s == 1 {
			return n
		}
	}
	return 0
}
