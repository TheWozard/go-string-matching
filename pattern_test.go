package matcher_test

import (
	"testing"

	matcher "github.com/TheWozard/go-string-matcher"
	"github.com/stretchr/testify/assert"
)

func TestPattern(t *testing.T) {
	basicPattern := matcher.CaseInsensitiveAndSymbolInsensitive.Matcher().Pattern("abc", "def")

	assert.True(t, basicPattern.Matches("abc"))
	assert.True(t, basicPattern.Matches("def"))
	assert.False(t, basicPattern.Matches("ghi"))
	assert.True(t, basicPattern.Matches("abc def"))

	withNegations := basicPattern.WithNegations("no", "non-")

	assert.True(t, withNegations.Matches("abc"))
	assert.False(t, withNegations.Matches("no abc"))
	assert.False(t, withNegations.Matches("no-abc"))
	assert.False(t, withNegations.Matches("noabc"))
	assert.False(t, withNegations.Matches("no abc no abc"))
	assert.True(t, withNegations.Matches("no abc no abc abc"))

	assert.False(t, withNegations.Matches("non-abc"))
	assert.False(t, withNegations.Matches("non abc"))
	assert.True(t, withNegations.Matches("nonabc"))
}
