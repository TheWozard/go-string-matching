package matcher_test

import (
	"fmt"
	"testing"

	matcher "github.com/TheWozard/go-string-matcher"
	"github.com/stretchr/testify/assert"
)

func TestASCIIBaseline(t *testing.T) {
	for i := ' '; i <= '~'; i++ {
		assert.Equal(t, i, matcher.ASCIIBaseline.Lookup(i).Equality, fmt.Sprintf("unexpected character for %q", i))
	}
}
