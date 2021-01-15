// Made with ❤️ by Dmytro Vovk for Geckoboard
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dmitry-vovk/trigram/trigrammer"
)

func main() {
	http.HandleFunc("/learn", learnHandler)
	http.HandleFunc("/generate", generateHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// learnHandler handles POST requests to '/learn'
func learnHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Invalid method %q, expected POST", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if ct := r.Header.Get("Content-Type"); ct != "text/plain" {
		log.Printf("Invalid content type %q, expected text/plain", ct)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print("Learning...")
	start := time.Now()
	trigrammer.Learn(r.Body)
	log.Printf("Learned in %s", time.Since(start))
	_ = r.Body.Close()
}

// generateHandler handles GET requests to '/generate'
func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid method %q, expected GET", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	word1 := r.URL.Query().Get("word1")
	if word1 == "" {
		log.Printf("Expected non-empty word1 parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	word2 := r.URL.Query().Get("word2")
	if word2 == "" {
		log.Printf("Expected non-empty word2 parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Generating using words %q & %q...", word1, word2)
	start := time.Now()
	text := trigrammer.Generate(word1, word2)
	log.Printf("Generated in %s", time.Since(start))
	w.Header().Add("Content-type", "text/plain")
	if _, err := w.Write([]byte(text)); err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
