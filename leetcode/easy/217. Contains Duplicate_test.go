package easy

import (
	"fmt"
	"testing"
)

func Test_217(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]interface{}, len(nums))
	for i, n := range nums {
		if m[n] == nil {
			m[n] = i
		} else {
			if i-m[n].(int) <= k {
				return true
			} else {
				m[n] = i
			}
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]interface{}, len(nums))
	for _, n := range nums {
		if m[n] != nil {
			return true
		}
		m[n] = n
	}
	return false
}
