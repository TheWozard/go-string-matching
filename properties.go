package matcher

var (
	ASCIIBaseline = ASCIIOffsetLookup{
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
	}

	ASCIICaseInsensitive = ASCIIBaseline.Copy().SetEquality(map[rune]rune{
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
)

type CharacterProperties struct {
	Equality  rune
	Ignorable bool
}

type ASCIIOffsetLookup []CharacterProperties

func (alm ASCIIOffsetLookup) Matcher() Matcher {
	return Matcher{Lookup: alm}
}

func (alm ASCIIOffsetLookup) Lookup(value rune) CharacterProperties {
	if value > '~' {
		return CharacterProperties{Equality: value, Ignorable: false}
	}
	return alm[value]
}

func (alm ASCIIOffsetLookup) Copy() ASCIIOffsetLookup {
	result := make(ASCIIOffsetLookup, len(alm))
	copy(result, alm)
	return result
}

func (alm ASCIIOffsetLookup) SetEquality(updates map[rune]rune) ASCIIOffsetLookup {
	for base, value := range updates {
		alm[base].Equality = value
	}
	return alm
}

func (alm ASCIIOffsetLookup) SetIgnorable(updates []rune) ASCIIOffsetLookup {
	for base := range updates {
		alm[base].Ignorable = true
	}
	return alm
}

func (alm ASCIIOffsetLookup) SetNotIgnorable(updates []rune) ASCIIOffsetLookup {
	for base := range updates {
		alm[base].Ignorable = false
	}
	return alm
}
