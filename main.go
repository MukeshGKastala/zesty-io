package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mkastala/zesty-io/trie"
)

var (
	// Command-Line Flags
	file, port string
	limit      uint
)

func init() {
	flag.StringVar(&file, "f", "./data/shakespeare-sample.txt", "path to data source")
	flag.UintVar(&limit, "l", 25, "limit the number of results")
	flag.StringVar(&port, "p", "9000", "port the HTTP server listens on")
}

func generatePrefixTree(file string) *trie.Trie {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	prefixTree := trie.New()

	for scanner.Scan() {
		// Consciously ignoring errors
		// TODO(MukeshKastala): Output failed inserts to file for inspection
		prefixTree.Insert(scanner.Text())
	}

	return prefixTree
}

func autocompleteHandler(pt *trie.Trie) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			r.ParseForm()

			params, ok := r.Form["prefix"]
			if !ok {
				http.Error(w, "400 missing query parameter", 400)
				return
			}
			if len(params) != 1 {
				http.Error(w, "422 malformatted query parameter", 422)
				return
			}

			prefix := params[0]

			words := struct {
				Words []string `json:"words"`
			}{
				Words: pt.Autocomplete(prefix, limit),
			}
			json.NewEncoder(w).Encode(words)
		default:
			http.Error(w, "405 GET is the only supported method", 405)
		}
	}
	return http.HandlerFunc(fn)
}

func main() {
	flag.Parse()

	// Locally-scoped ServeMux for added security
	// TODO(MukeshKastala): Add logging and error handling middleware - This is made easy using a HTTP web framework like gin
	mux := http.NewServeMux()

	pt := generatePrefixTree(file)
	mux.Handle("/autocomplete", autocompleteHandler(pt))

	log.Println("HTTP server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
