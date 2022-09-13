package service

func (a *Artist) Draw() {
	var notunique bool
	a.Art = ""
	for ix, line := range a.Text {
		if line == "" {
			if !notunique && ix == len(a.Text)-1 {
				break
			}
			continue
		}
		notunique = true
		for i := 0; i < 8; i++ {
			for _, ch := range line {
				a.Art += string(a.Template[ch][i])
			}
			a.Art += "\n"
		}
	}

}
