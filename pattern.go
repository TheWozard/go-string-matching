package matcher

type Pattern struct {
	Words     []string
	Negations []string
	// Exclusions []string
	Matcher Matcher
}

func (p Pattern) WithNegations(words ...string) Pattern {
	return Pattern{Words: p.Words, Negations: append(p.Negations, words...), Matcher: p.Matcher}
}

func (p Pattern) WithWords(words ...string) Pattern {
	return Pattern{Words: append(p.Words, words...), Negations: p.Negations, Matcher: p.Matcher}
}

func (p Pattern) Matches(value string) bool {
Words:
	for _, word := range p.Words {
		index := 0
	Search:
		for index >= 0 && index <= len(value)-len(word) {
			start, end := p.Matcher.Index(word, value[index:])
			// It has found something, so we need to check if it is negated.
			if start == -1 || end == -1 {
				continue Words
			}
			for _, negation := range p.Negations {
				if p.Matcher.HasSuffix(negation, value[:index+start]) {
					index += end
					continue Search
				}
			}
			return true
		}
	}
	return false
}
