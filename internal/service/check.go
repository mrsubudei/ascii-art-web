package service

import (
	"crypto/sha256"
	"fmt"
	"unicode"
)

func isASCII(text string) string {
	chars := ""
	for ix, a := range text {
		if a > unicode.MaxASCII {
			chars += string(a)
			if len(text)-1 > ix {
				chars += ","
			}
		}
	}
	if chars != "" {
		return fmt.Sprintf("Not ASCII character(s) (%s)\n", chars)
	}
	return ""
}

func validateTemplate(contents, banner string) bool {
	hashs := map[string]string{
		"shadow":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"thinkertoy": "a57beec43fde6751ba1d30495b092658a064452f321e221d08c3ac34a9dc1294",
		"standard":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents))) == hashs[banner]
}

func (a *Artist) HasError() bool {
	if a.ErrMessage != "" {
		return true
	}
	return false
}
