package ascii

import (
	"os"
	"strings"
)

func ConvertAscii(arg, font string) (string, int) {
	file1, err := os.Open("pkg/ascii/fonts/" + font + ".txt")
	if err != nil {
		return "Internal Server Error. Font file does not exist", 500
	}

	if !TxtFileCheck(font + ".txt") {
		return "Internal Server Error. Font file was damaged", 500
	}

	words := strings.Split(arg, "\r\n")

	ascii := MakeMap(file1)
	if !CheckRange(words, ascii) {
		return "Bad request. This kind of symbols are not Allowed", 400
	}
	return PrintAns(words, ascii), 200
}
