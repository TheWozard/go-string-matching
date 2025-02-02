package matcher

import "unicode/utf8"

type Matcher struct {
	Lookup PropertiesLookup
}

// Matches returns true if the pattern is equal to the value.
func (m Matcher) Matches(pattern, value string) bool {
	ok, end := m.endIndex(pattern, value)
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
	startProperties, _ := m.Lookup.LookupNextRune(pattern, 0)
	i := 0
	for i < len(value) {
		valueProperties, result := m.Lookup.LookupNextRune(value, i)
		if valueProperties.Equality == startProperties.Equality {
			ok, _ := m.endIndex(pattern, value[i:])
			if ok {
				return true
			}
		}
		i = result
	}
	return false
}

// endIndex returns the index of the last byte in value that matches the pattern.
func (m Matcher) endIndex(pattern, value string) (bool, int) {
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
