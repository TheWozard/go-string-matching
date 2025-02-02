package matcher_test

import (
	"fmt"
	"testing"

	matcher "github.com/TheWozard/go-string-matcher"
	"github.com/stretchr/testify/assert"
)

func TestLookupNextRune(t *testing.T) {
	testCases := []struct {
		desc          string
		lookup        matcher.PropertiesLookup
		input         string
		index         int
		expected      matcher.Properties
		expectedIndex int
	}{
		// {
		// 	desc:   "empty string",
		// 	lookup: matcher.Baseline,
		// 	input:  "",
		// 	index:  0,
		// 	expected: matcher.Properties{
		// 		Equality:  -1,
		// 		Ignorable: true,
		// 	},
		// 	expectedIndex: 0,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, actualIndex := tC.lookup.LookupNextRune(tC.input, tC.index)
			assert.Equal(t, tC.expected, actual)
			assert.Equal(t, tC.expectedIndex, actualIndex)
		})
	}
}

func TestBaseline(t *testing.T) {
	for i := ' '; i <= '~'; i++ {
		assert.Equal(t, i, matcher.Baseline.LookupRune(i).Equality, fmt.Sprintf("unexpected character for %q", i))
	}
}
