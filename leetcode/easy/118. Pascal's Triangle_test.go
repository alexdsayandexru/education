package easy

import (
	"fmt"
	"testing"
)

func Test_118(t *testing.T) {
	fmt.Println(getRow(6))

}
func getRow(rowIndex int) []int {
	return generate(rowIndex + 1)[rowIndex]
}

func generate(numRows int) [][]int {
	var out [][]int
	var prev []int
	for i := range numRows {
		if i == 0 {
			out = append(out, []int{1})
		} else if i == 1 {
			prev = []int{1, 1}
			out = append(out, prev)
		} else {
			curr := make([]int, len(prev)+1)
			for i := range len(prev) {
				if i-1 < 0 {
					curr[i] = 1
				} else {
					curr[i] = prev[i] + prev[i-1]
				}
				if i == len(prev)-1 {
					curr[i+1] = 1
				}
			}
			out = append(out, curr)
			prev = curr
		}
	}
	return out
}
