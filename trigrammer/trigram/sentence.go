package trigram

import "strings"

// Sentence holds words from trigrams to create a text sequence
type Sentence struct {
	words []string
}

// NewSentence returns new Sentence instance based on the seed trigram
func NewSentence(tr *Trigram) *Sentence {
	return &Sentence{
		words: []string{
			tr.Word1,
			tr.Word2,
			tr.Word3,
		},
	}
}

// Add appends another trigram to the sentence
func (s *Sentence) Add(tr *Trigram) *Sentence {
	s.words = append(s.words, tr.Word3)
	return s
}

// String implements io.Stringer
func (s Sentence) String() string { return strings.Join(s.words, " ") }
