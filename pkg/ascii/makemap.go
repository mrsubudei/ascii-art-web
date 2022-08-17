package ascii

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
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

func TxtFileCheck(fileName string) bool {
	hashStandard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	hashShadow := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	hashThinkertoy := "a57beec43fde6751ba1d30495b092658a064452f321e221d08c3ac34a9dc1294"

	file, err := os.Open("pkg/ascii/fonts/" + fileName)
	if err != nil {
		log.Println(err)
		return false
	}
	defer file.Close()
	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
	sum := fmt.Sprintf("%x", sha256.Sum(nil))
	switch fileName {
	case "shadow.txt":
		if string(sum) == string(hashShadow) {
			return true
		}
	case "standard.txt":
		if string(sum) == string(hashStandard) {
			return true
		}
	case "thinkertoy.txt":
		if string(sum) == string(hashThinkertoy) {
			return true
		}
	}
	return false
}
