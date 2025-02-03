package matcher

import "unicode/utf8"

type Matcher struct {
	Lookup PropertiesLookup
}

func (m Matcher) Pattern(words ...string) Pattern {
	return Pattern{Matcher: m, Words: words}
}

// Matches returns true if the pattern is equal to the value.
func (m Matcher) Matches(pattern, value string) bool {
	ok, end := m.prefixIndex(pattern, value)
	if !ok {
		return false
	}
	var valueProperties Properties
	for end < len(value) {
		valueProperties, end = m.Lookup.LookupNextRune(value, end)
		if !valueProperties.Ignorable {
			return false
		}
	}
	return true
}

// Contains returns true if the pattern is contained within the value.
func (m Matcher) Contains(pattern, value string) bool {
	start, _ := m.Index(pattern, value)
	return start != -1
}

// Index returns the start and end index of the pattern in the value.
func (m Matcher) Index(pattern, value string) (int, int) {
	if len(value) < len(pattern) {
		return -1, -1
	}
	startProperties, _ := m.Lookup.LookupNextRune(pattern, 0)
	i := 0
	for i < len(value) {
		valueProperties, result := m.Lookup.LookupNextRune(value, i)
		if valueProperties.Equality == startProperties.Equality {
			ok, end := m.prefixIndex(pattern, value[i:])
			if ok {
				return i, i + end
			}
		}
		i = result
	}
	return -1, -1
}

// HasPrefix returns true if the pattern is a prefix of the value.
func (m Matcher) HasPrefix(pattern, value string) bool {
	ok, _ := m.prefixIndex(pattern, value)
	return ok
}

func (m Matcher) HasSuffix(pattern, value string) bool {
	ok, _ := m.suffixIndex(pattern, value)
	return ok
}

// prefixIndex returns the index of the last byte in value that matches the pattern.
func (m Matcher) prefixIndex(pattern, value string) (bool, int) {
	// If the value is shorter than the pattern, then it cannot contain the pattern.
	if len(value) < len(pattern) {
		return false, 0
	}
	var patternProperties, valueProperties Properties
	i, j := 0, 0
PATTERN:
	for i < len(pattern) {
		if pattern[i] < utf8.RuneSelf {
			patternProperties = m.Lookup.ASCII[pattern[i]]
			i++
		} else {
			patternProperties, i = m.Lookup.LookupNextRune(pattern, i)
		}
		for j < len(value) {
			if value[j] < utf8.RuneSelf {
				valueProperties = m.Lookup.ASCII[value[j]]
				j++
			} else {
				valueProperties, j = m.Lookup.LookupNextRune(value, j)
			}
			if valueProperties.Equality == patternProperties.Equality {
				continue PATTERN
			}
			if valueProperties.Ignorable {
				continue
			}
			return false, 0
		}
		return false, 0
	}
	return true, j
}

// suffixIndex returns the index of the first byte in value that matches the pattern.
func (m Matcher) suffixIndex(pattern, value string) (bool, int) {
	// If the value is shorter than the pattern, then it cannot contain the pattern.
	if len(value) < len(pattern) {
		return false, 0
	}
	var patternProperties, valueProperties Properties
	i, j := len(pattern)-1, len(value)-1
PATTERN:
	for i > 0 {
		if pattern[i] < utf8.RuneSelf {
			patternProperties = m.Lookup.ASCII[pattern[i]]
			i--
		} else {
			patternProperties, i = m.Lookup.LookupPreviousRune(pattern, i)
		}
		for j > 0 {
			if value[j] < utf8.RuneSelf {
				valueProperties = m.Lookup.ASCII[value[j]]
				j--
			} else {
				valueProperties, j = m.Lookup.LookupPreviousRune(value, j)
			}
			if valueProperties.Equality == patternProperties.Equality {
				continue PATTERN
			}
			if valueProperties.Ignorable {
				continue
			}
			return false, 0
		}
		return false, 0
	}
	return true, j
}
