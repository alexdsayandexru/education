package easy

import (
	"fmt"
	"testing"
)

func Test_168(t *testing.T) {
	fmt.Println(convertToTitle(19010))
}

func convertToTitle(columnNumber int) string {
	var result []byte

	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		result = append(result, byte('A'+remainder))
		columnNumber /= 26
	}

	// Разворачиваем срез, так как буквы были добавлены в обратном порядке
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
