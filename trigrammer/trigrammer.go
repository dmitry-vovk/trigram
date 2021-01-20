package trigrammer

import (
	"bufio"
	"io"

	"github.com/dmitry-vovk/trigram/trigrammer/storage"
	"github.com/dmitry-vovk/trigram/trigrammer/trigram"
)

// Generate returns generated text starting with the seed word1 & word2
func Generate(word1, word2 string) string {
	tr := storage.GetMatchingTrigram(&trigram.Trigram{Word2: word1, Word3: word2})
	if tr == nil {
		return ""
	}
	sentence := trigram.NewSentence(tr)
	for {
		if tr = storage.GetMatchingTrigram(tr); tr == nil {
			break
		}
		sentence.Add(tr)
	}
	return sentence.String()
}

// Learn consumes text to be split into trigrams
func Learn(text io.Reader) error {
	done := make(chan error, 1)
	for t := range extractTrigrams(text, done) {
		storage.AddTrigram(t)
	}
	return <-done
}

// extractTrigrams scans and emits trigrams from the input reader
func extractTrigrams(input io.Reader, done chan error) chan *trigram.Trigram {
	var words []string
	trigramC := make(chan *trigram.Trigram)
	go func(c chan *trigram.Trigram) {
		for word := range scanTrigrams(input, done) {
			words = append(words, word)
			if len(words) == 3 {
				trigramC <- &trigram.Trigram{Word1: words[0], Word2: words[1], Word3: words[2]}
				words[0], words[1] = words[1], words[2]
				words = words[0:2]
			}
		}
		close(c)
	}(trigramC)
	return trigramC
}

// scanTrigrams scans and emits words from the input reader
func scanTrigrams(input io.Reader, done chan error) chan string {
	outC := make(chan string)
	go func(c chan string) {
		r := bufio.NewScanner(input)
		r.Split(splitter)
		for r.Scan() {
			c <- r.Text()
		}
		if err := r.Err(); err != nil {
			done <- err
		}
		close(c)
		close(done)
	}(outC)
	return outC
}

// splitter defines how words are extracted from input data
func splitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Can use another splitter here
	// e.g. to treat punctuation differently
	return bufio.ScanWords(data, atEOF)
}
