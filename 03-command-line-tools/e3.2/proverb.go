package main

type proverb struct {
	line  string
	chars map[rune]int
}

func (p *proverb) countChars() {
	if p.chars != nil {
		return
	}

	m := make(map[rune]int, 0)
	for _, c := range p.line {
		m[c] = m[c] + 1
	}
	p.chars = m
}

func newProverb(line string) *proverb {
	p := new(proverb)
	p.line = line
	p.countChars()
	return p
}
