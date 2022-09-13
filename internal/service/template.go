package service

import (
	"bufio"
	"fmt"
	"os"
)

func Construct(banner string) (map[rune][8]string, string) {
	filename := fmt.Sprintf("templates/fonts/%s.txt", banner)
	errMessage := "Template may have been corrupted"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, errMessage
	}
	readFile := bufio.NewReader(file)
	contents := ""
	template := map[rune][8]string{}
	for i := rune(32); i < 127; i++ {
		s, _, err := readFile.ReadLine()
		if err != nil {
			return nil, errMessage
		}
		contents += string(s) + "\n"
		var letter [8]string
		for j := 0; j < 8; j++ {
			l, _, err := readFile.ReadLine()
			if err != nil {
				return nil, errMessage
			}
			line := string(l)
			contents += line + "\n"
			letter[j] = line
		}
		template[i] = letter
	}
	if !validateTemplate(contents, banner) {
		return nil, errMessage
	}
	return template, ""
}
