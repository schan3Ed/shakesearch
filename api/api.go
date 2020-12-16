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

}

func HandleSearch(loader model.FileLoadSharder, searcher model.Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}

		var results []string
		for _, f := range loader.Shards {
			res := searcher.Search(query[0], f.CompleteWorkShard, f.SuffixShard)
			results = append(results, res...)
		}
		sendJson(w, results)
	}
}

