package med

import "testing"

func Test_12(t *testing.T) {
	if intToRoman(3749) != "MMMDCCXLIX" {
		t.Error()
	}
}

func intToRoman(num int) string {
	m := getRomanMap()
	snum := ""

	for _, i := range []int{1000, 100, 10, 1} {
		n := num / i
		if n != 0 {
			snum = snum + m[n*i]
			num = num % i
		}
	}

	return snum
}

func getRomanMap() map[int]string {
	m := make(map[int]string, 42)

	m[1] = "I"
	m[2] = "II"
	m[3] = "III"
	m[4] = "IV"
	m[5] = "V"
	m[6] = "VI"
	m[7] = "VII"
	m[8] = "VIII"
	m[9] = "IX"
	m[10] = "X"
	m[20] = "XX"
	m[30] = "XXX"
	m[40] = "XL"
	m[50] = "L"
	m[60] = "LX"
	m[70] = "LXX"
	m[80] = "LXXX"
	m[90] = "XC"
	m[100] = "C"
	m[200] = "CC"
	m[300] = "CCC"
	m[400] = "CD"
	m[500] = "D"
	m[600] = "DC"
	m[700] = "DCC"
	m[800] = "DCCC"
	m[900] = "CM"
	m[1000] = "M"
	m[2000] = "MM"
	m[3000] = "MMM"

	return m
}
