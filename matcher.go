package matcher

type Matcher struct {
	Lookup ASCIIOffsetLookup
}

func (m Matcher) Matches(pattern, value string) bool {
	i, j := 0, 0
	for ; i < len(pattern); i++ {
		patternCharProperties := m.Lookup.Lookup(rune(pattern[i]))
		for {
			if i+j >= len(value) {
				return false
			}
			valueCharProperties := m.Lookup.Lookup(rune(value[i+j]))
			if valueCharProperties.Equality == patternCharProperties.Equality {
				break
			}
			if valueCharProperties.Ignorable {
				j++
				continue
			}
			return false
		}
	}
	if i+j < len(value) {
		for ; i+j < len(value); j++ {
			valueCharProperties := m.Lookup.Lookup(rune(value[i+j]))
			if !valueCharProperties.Ignorable {
				return false
			}
		}
	}
	return true
}

func (m Matcher) Contains(a, b string) bool {
	return false
}
