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
