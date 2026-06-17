package easy

import (
	"fmt"
	"testing"
)

func Test_169(t *testing.T) {
	nums := []int{3, 2, 2, 3, 1, 5, 1, 1, 2, 7, 2, 8, 9, 10}
	fmt.Println(majorityElement2(nums))
}

func majorityElement2(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	r := 0
	mx := 0

	for n, s := range m {
		if s > mx {
			r = n
			mx = s
		}
	}
	return r
}
