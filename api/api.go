package api

import (
	"bytes"
	"encoding/json"
	"net/http"

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

func HandleSearch(loader model.FileLoadSharder, searcher model.Searcher) func(w http.ResponseWriter, r *http.Request) {
	var s SearchResult
	s.Results = make(map[string][]string)
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}

		results := make(map[string][]string)
		var total int
		for _, f := range loader.Shards {
			res := searcher.Search(query[0], f.CompleteWorkShard, f.SuffixShard)
			if len(res) == 0 {
				continue
			}
			results[f.Title] = res
			total += len(res)
		}
		s.Results = results
		s.Occur = total
		sendJson(w, s)
	}
}

