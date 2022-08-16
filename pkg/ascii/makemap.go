package ascii

import (
	"bufio"
	"os"
)

func MakeMap(file *os.File) map[rune][9]string {
	scanner := bufio.NewScanner(file)
	ans := make(map[rune][9]string)
	slice := rune(32)
	ind := 0
	var tmp [9]string

	for scanner.Scan() {
		tmp[ind] = scanner.Text()
		ind++
		if ind == 9 {
			ans[slice] = tmp
			ind = 0
			slice++
		}
	}
	return ans
}

func CheckRange(words []string, ascii map[rune][9]string) bool {
	for _, val := range words {
		for _, v := range val {
			if !(v >= 32 && v <= 126) {
				return false
			}
		}
	}
	return true
}

func PrintAns(words []string, ascii map[rune][9]string) string {
	answer := ""

	for _, val := range words {
		if val == "" {
			answer += " "
			answer += "\n"
			continue
		}
		for i := 1; i < 9; i++ {

			for _, v := range val {
				answer += ascii[v][i]
			}
			answer += "\n"

		}
	}
	return answer
}
