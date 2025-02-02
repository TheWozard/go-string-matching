package matcher_test

import (
	"regexp"
	"strings"
	"testing"

	matcher "github.com/TheWozard/go-string-matcher"
	"github.com/stretchr/testify/assert"
)

func BenchmarkEquality(b *testing.B) {
	pattern, value := "lorem ipsum dolor sit amet", "LOREM IPSUM DOLOR SIT AMET"

	// Fastest, but does not support all features needed.
	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, strings.EqualFold(pattern, value))
		}
	})

	// Ours
	b.Run("Matcher", func(b *testing.B) {
		CaseInsensitive := matcher.CaseInsensitive.Matcher()
		for i := 0; i < b.N; i++ {
			assert.True(b, CaseInsensitive.Matches(pattern, value))
		}
	})

	// Slower common naive approach. Does not support all features needed.
	b.Run("ToLower", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			assert.True(b, pattern == strings.ToLower(value))
		}
	})

	// Slowest, but supports all features needed.
	b.Run("Regex", func(b *testing.B) {
		matcher := regexp.MustCompile("(?i)lorem ipsum dolor sit amet")
		for i := 0; i < b.N; i++ {
			assert.True(b, matcher.MatchString(value))
		}
	})
}

func TestMatcherMatches(t *testing.T) {
	CaseInsensitive := matcher.CaseInsensitive.Matcher()

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

	// Unicode
	assert.True(t, CaseInsensitive.Matches("👍", "👍"))
	assert.True(t, CaseInsensitive.Matches("👍 Nice Work", "👍 nice work\n"))
	assert.False(t, CaseInsensitive.Matches("👍", "👎"))
}

func TestMatcherContains(t *testing.T) {
	CaseInsensitive := matcher.CaseInsensitive.Matcher()

	// Word Only
	assert.True(t, CaseInsensitive.Contains("abc", "ABC"))
	assert.True(t, CaseInsensitive.Contains("abc", "abc"))
	assert.True(t, CaseInsensitive.Contains("abc", "aBc"))

	// Suffix Variations
	assert.False(t, CaseInsensitive.Contains("abc", "ab"))
	assert.True(t, CaseInsensitive.Contains("abc", "abcd"))
	assert.True(t, CaseInsensitive.Contains("abc", "abc "))
	assert.True(t, CaseInsensitive.Contains("abc", "abc\n"))

	// Prefix Variations
	assert.False(t, CaseInsensitive.Contains("abc", "bc"))
	assert.True(t, CaseInsensitive.Contains("abc", "dabc"))
	assert.True(t, CaseInsensitive.Contains("abc", " abc"))
	assert.True(t, CaseInsensitive.Contains("abc", "\nabc"))

	// Internal Variations
	assert.False(t, CaseInsensitive.Contains("abc", "ac"))
	assert.False(t, CaseInsensitive.Contains("abc", "abdc"))
	assert.False(t, CaseInsensitive.Contains("abc", "ab c"))
	assert.True(t, CaseInsensitive.Contains("abc", "ab\nc"))

	// Unicode
	assert.True(t, CaseInsensitive.Contains("👍", "👍"))
	assert.True(t, CaseInsensitive.Contains("👍 Nice Work", "That is some really 👍 nice work. Good job."))
	assert.False(t, CaseInsensitive.Contains("👍", "👎"))
}
