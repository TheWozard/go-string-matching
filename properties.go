package matcher

import (
	"unicode/utf8"
)

var (
	Baseline = PropertiesLookup{
		ASCII: []Properties{
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' ', Ignorable: true},
			{Equality: ' '},
			{Equality: '!'},
			{Equality: '"'},
			{Equality: '#'},
			{Equality: '$'},
			{Equality: '%'},
			{Equality: '&'},
			{Equality: '\''},
			{Equality: '('},
			{Equality: ')'},
			{Equality: '*'},
			{Equality: '+'},
			{Equality: ','},
			{Equality: '-'},
			{Equality: '.'},
			{Equality: '/'},
			{Equality: '0'},
			{Equality: '1'},
			{Equality: '2'},
			{Equality: '3'},
			{Equality: '4'},
			{Equality: '5'},
			{Equality: '6'},
			{Equality: '7'},
			{Equality: '8'},
			{Equality: '9'},
			{Equality: ':'},
			{Equality: ';'},
			{Equality: '<'},
			{Equality: '='},
			{Equality: '>'},
			{Equality: '?'},
			{Equality: '@'},
			{Equality: 'A'},
			{Equality: 'B'},
			{Equality: 'C'},
			{Equality: 'D'},
			{Equality: 'E'},
			{Equality: 'F'},
			{Equality: 'G'},
			{Equality: 'H'},
			{Equality: 'I'},
			{Equality: 'J'},
			{Equality: 'K'},
			{Equality: 'L'},
			{Equality: 'M'},
			{Equality: 'N'},
			{Equality: 'O'},
			{Equality: 'P'},
			{Equality: 'Q'},
			{Equality: 'R'},
			{Equality: 'S'},
			{Equality: 'T'},
			{Equality: 'U'},
			{Equality: 'V'},
			{Equality: 'W'},
			{Equality: 'X'},
			{Equality: 'Y'},
			{Equality: 'Z'},
			{Equality: '['},
			{Equality: '\\'},
			{Equality: ']'},
			{Equality: '^'},
			{Equality: '_'},
			{Equality: '`'},
			{Equality: 'a'},
			{Equality: 'b'},
			{Equality: 'c'},
			{Equality: 'd'},
			{Equality: 'e'},
			{Equality: 'f'},
			{Equality: 'g'},
			{Equality: 'h'},
			{Equality: 'i'},
			{Equality: 'j'},
			{Equality: 'k'},
			{Equality: 'l'},
			{Equality: 'm'},
			{Equality: 'n'},
			{Equality: 'o'},
			{Equality: 'p'},
			{Equality: 'q'},
			{Equality: 'r'},
			{Equality: 's'},
			{Equality: 't'},
			{Equality: 'u'},
			{Equality: 'v'},
			{Equality: 'w'},
			{Equality: 'x'},
			{Equality: 'y'},
			{Equality: 'z'},
			{Equality: '{'},
			{Equality: '|'},
			{Equality: '}'},
			{Equality: '~'},
		},
		Unicode: map[rune]Properties{},
	}

	CaseInsensitive = Baseline.Copy().SetEquality(map[rune]rune{
		'A': 'a',
		'B': 'b',
		'C': 'c',
		'D': 'd',
		'E': 'e',
		'F': 'f',
		'G': 'g',
		'H': 'h',
		'I': 'i',
		'J': 'j',
		'K': 'k',
		'L': 'l',
		'M': 'm',
		'N': 'n',
		'O': 'o',
		'P': 'p',
		'Q': 'q',
		'R': 'r',
		'S': 's',
		'T': 't',
		'U': 'u',
		'V': 'v',
		'W': 'w',
		'X': 'x',
		'Y': 'y',
		'Z': 'z',
	})

	CaseInsensitiveAndSymbolInsensitive = CaseInsensitive.Copy().SetIgnorable([]rune{
		' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~',
	}, true)
)

// Properties defines how a given rune should be treated.
type Properties struct {
	Equality  rune
	Ignorable bool
}

// PropertiesLookup supports methods for looking up properties and defining properties of runes.
type PropertiesLookup struct {
	ASCII   []Properties
	Unicode map[rune]Properties
}

// Matcher converts this PropertiesLookup into a Matcher.
func (alm PropertiesLookup) Matcher() Matcher {
	return Matcher{Lookup: alm}
}

func (alm PropertiesLookup) LookupRune(value rune) Properties {
	if value < utf8.RuneSelf {
		return alm.ASCII[value]
	}
	if properties, ok := alm.Unicode[value]; ok {
		return properties
	}
	return Properties{Equality: value, Ignorable: false}
}

func (alm PropertiesLookup) LookupNextRune(str string, index int) (Properties, int) {
	value, size := utf8.DecodeRuneInString(str[index:])
	if size == 0 {
		return Properties{Equality: -1, Ignorable: true}, index
	}
	return alm.LookupRune(value), size + index
}

func (alm PropertiesLookup) Copy() PropertiesLookup {
	result := PropertiesLookup{
		ASCII:   make([]Properties, len(alm.ASCII)),
		Unicode: make(map[rune]Properties, len(alm.Unicode)),
	}
	copy(result.ASCII, alm.ASCII)
	for key, value := range alm.Unicode {
		result.Unicode[key] = value
	}
	return result
}

func (alm PropertiesLookup) SetEquality(updates map[rune]rune) PropertiesLookup {
	for base, value := range updates {
		if base < utf8.RuneSelf {
			alm.ASCII[base].Equality = value
		} else {
			prev := alm.Unicode[base]
			prev.Equality = value
			alm.Unicode[base] = prev
		}
	}
	return alm
}

func (alm PropertiesLookup) SetIgnorable(updates []rune, value bool) PropertiesLookup {
	for _, base := range updates {
		if base < utf8.RuneSelf {
			alm.ASCII[base].Ignorable = true
		} else {
			prev, ok := alm.Unicode[base]
			if !ok {
				prev = Properties{Equality: base}
			}
			prev.Ignorable = value
			alm.Unicode[base] = prev
		}
	}
	return alm
}
