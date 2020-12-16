package model

import (
	"index/suffixarray"
)

//type Searcher interface {
//	ExactMatch(s string)
//	FuzzyLevenshtein(s string)
//}


type Searcher struct {}

func findPrevious(works string, idx int) int {
	var count int
	for i := idx; i != 0; i-- {
		if works[i] == '\n' {
			count++
			if count == 4 {
				return i
			}
		}
	}
	return 0
}

func findAfter(works string, idx int) int {
	var count int
	for i := idx; i != len(works) - 1; i++ {
		if works[i] == '\n' {
			count++
			if count == 4 {
				return i
			}
		}
	}
	return len(works) - 1
}

func (s *Searcher) Search(query string, works string, index *suffixarray.Index) []string {
	idxs := index.Lookup([]byte(query), -1)
	results := []string{}
	for _, idx := range idxs {
		results = append(results, works[findPrevious(works, idx):findAfter(works, idx)])
	}
	return results
}

