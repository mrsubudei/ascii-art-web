package service

import (
	"strings"
)

type Artist struct {
	Art        string
	Text       []string
	ErrMessage string
	StatusCode int
	Template   map[rune][8]string
	banners    []Banner
}

type Banner struct {
	name     string
	template map[rune][8]string
}

func (a *Artist) UpdateText(text string) {
	if text == "" {
		a.ErrMessage = "Empty Input."
		return
	}
	err := isASCII(text)
	if err != "" {
		a.ErrMessage = err
		return
	}
	a.Text = strings.Split(strings.Replace(text, "\\n", "\n", -1), "\n")
}

func (a *Artist) UpdateBanner(banner string) {
	for _, b := range a.banners {
		if b.name == banner {
			a.Template = b.template
			return
		}
	}
	newTemplate, err := Construct(banner)
	if err != "" {
		a.ErrMessage = err
		return
	}
	a.Template = newTemplate
	a.banners = append(a.banners, Banner{banner, a.Template})
}

func (a *Artist) UpdateArtist() {
	a.ErrMessage = ""
	a.Art = ""
}
