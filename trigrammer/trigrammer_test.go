package trigrammer

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/dmitry-vovk/trigram/trigrammer/trigram"
	"github.com/stretchr/testify/assert"
)

func TestSplitter(t *testing.T) {
	testCases := []struct {
		input    io.Reader
		expected []*trigram.Trigram
	}{
		{
			input: strings.NewReader(`
				It is a truth universally acknowledged,
				that a single man in possession of a good fortune,
				must be in want of a wife.`,
			),
			expected: []*trigram.Trigram{
				{`It`, `is`, `a`},
				{`is`, `a`, `truth`},
				{`a`, `truth`, `universally`},
				{`truth`, `universally`, `acknowledged,`},
				{`universally`, `acknowledged,`, `that`},
				{`acknowledged,`, `that`, `a`},
				{`that`, `a`, `single`},
				{`a`, `single`, `man`},
				{`single`, `man`, `in`},
				{`man`, `in`, `possession`},
				{`in`, `possession`, `of`},
				{`possession`, `of`, `a`},
				{`of`, `a`, `good`},
				{`a`, `good`, `fortune,`},
				{`good`, `fortune,`, `must`},
				{`fortune,`, `must`, `be`},
				{`must`, `be`, `in`},
				{`be`, `in`, `want`},
				{`in`, `want`, `of`},
				{`want`, `of`, `a`},
				{`of`, `a`, `wife.`},
			},
		},
		{
			input:    strings.NewReader(""),
			expected: nil,
		},
		{
			input:    strings.NewReader("two words"),
			expected: nil,
		},
		{
			input: strings.NewReader("whole three words"),
			expected: []*trigram.Trigram{
				{`whole`, `three`, `words`},
			},
		},
	}
	for _, tc := range testCases {
		var result []*trigram.Trigram
		done := make(chan error)
		for tr := range extractTrigrams(tc.input, done) {
			result = append(result, tr)
		}
		assert.Equal(t, tc.expected, result)
		assert.NoError(t, <-done)
	}
}

func TestAll(t *testing.T) {
	f, err := os.Open("test_data/pride-prejudice.txt")
	if err != nil {
		t.Fatalf("Error loading test file: %s", err)
	}
	defer func() { _ = f.Close() }()
	assert.NoError(t, Learn(f))
	assert.True(t, len(Generate("It", "is")) > 0)
}
