package matcher_test

import (
	"strings"
	"testing"

	matcher "github.com/TheWozard/go-string-matcher"
	"github.com/stretchr/testify/assert"
)

func BenchmarkEquality(b *testing.B) {
	pattern, value := "lorem ipsum dolor sit amet", "LOREM IPSUM DOLOR SIT AMET"

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, strings.EqualFold(pattern, value))
		}
	})

	b.Run("Matcher", func(b *testing.B) {
		CaseInsensitive := matcher.ASCIICaseInsensitive.Matcher()
		for i := 0; i < b.N; i++ {
			assert.True(b, CaseInsensitive.Matches(pattern, value))
		}
	})

	b.Run("ToLower", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, pattern == strings.ToLower(value))
		}
	})
}

func TestMatcher(t *testing.T) {
	CaseInsensitive := matcher.ASCIICaseInsensitive.Matcher()

	// Word Only
	assert.True(t, CaseInsensitive.Matches("abc", "ABC"))
	assert.True(t, CaseInsensitive.Matches("abc", "abc"))
	assert.True(t, CaseInsensitive.Matches("abc", "aBc"))

	// Suffix Variations
	assert.False(t, CaseInsensitive.Matches("abc", "ab"))
	assert.False(t, CaseInsensitive.Matches("abc", "abcd"))
	assert.False(t, CaseInsensitive.Matches("abc", "abc "))
	assert.True(t, CaseInsensitive.Matches("abc", "abc\n"))

	// Prefix Variations
	assert.False(t, CaseInsensitive.Matches("abc", "bc"))
	assert.False(t, CaseInsensitive.Matches("abc", "dabc"))
	assert.False(t, CaseInsensitive.Matches("abc", " abc"))
	assert.True(t, CaseInsensitive.Matches("abc", "\nabc"))

	// Internal Variations
	assert.False(t, CaseInsensitive.Matches("abc", "ac"))
	assert.False(t, CaseInsensitive.Matches("abc", "abdc"))
	assert.False(t, CaseInsensitive.Matches("abc", "ab c"))
	assert.True(t, CaseInsensitive.Matches("abc", "ab\nc"))
}
