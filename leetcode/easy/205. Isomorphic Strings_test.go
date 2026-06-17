package easy

import (
	"fmt"
	"testing"
)

func Test_205(t *testing.T) {
	fmt.Println(isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	mst := make(map[byte]byte)
	mts := make(map[byte]byte)

	bs := []byte(s)
	bt := []byte(t)
	if len(bs) != len(bt) {
		return false
	}

	for i := range len(bs) {
		if mst[bs[i]] == 0 && mts[bt[i]] == 0 {
			mst[bs[i]] = bt[i]
			mts[bt[i]] = bs[i]
		} else if mst[bs[i]] != bt[i] || mts[bt[i]] != bs[i] {
			return false
		}
	}
	return true
}
