package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"pulley.com/shakesearch/model"
)

func sendJson(w http.ResponseWriter, body interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("encoding failure"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}

type SearchResult struct {
	Occur int
	Results map[string][]string
}

func searchByString(loader model.FileLoadSharder, searcher model.Searcher, s string, caseSen string, book string) SearchResult {
	var res SearchResult
	results := make(map[string][]string)
	var total int
	for _, f := range loader.Shards {
		if book != "" && f.Title != book {
			continue
		}
		var res []string
		if caseSen == "true" {
			res = searcher.Search(strings.ToLower(s), f.CompleteWorkShard, f.CaseSuffixShard)
		} else {
			res = searcher.Search(s, f.CompleteWorkShard, f.SuffixShard)
		}
		if len(res) == 0 {
			continue
		}
		results[f.Title] = res
		total += len(res)
	}
	res.Results = results
	res.Occur = total
	return res
}

func HandleSearch(loader model.FileLoadSharder, searcher model.Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		caseSen, ok := r.URL.Query()["case"]
		bookName, ok := r.URL.Query()["name"]
		var cas, book string
		if len(bookName) == 1 {
			book = bookName[0]
		}
		if len(caseSen) == 1 {
			cas = caseSen[0]
		}
		sendJson(w, searchByString(loader, searcher, query[0],cas, book))
	}
}

