package model

import (
	"bufio"
	"fmt"
	"index/suffixarray"
	"os"
	"strings"
)

var bookTitles = map[string]bool{
	"THE SONNETS":                              true,
	"ALL’S WELL THAT ENDS WELL":                true,
	"THE TRAGEDY OF ANTONY AND CLEOPATRA":      true,
	"AS YOU LIKE IT":                           true,
	"THE COMEDY OF ERRORS": true,
	"THE TRAGEDY OF CORIOLANUS": true,
	"CYMBELINE": true,
	"THE TRAGEDY OF HAMLET, PRINCE OF DENMARK": true,
	"THE FIRST PART OF KING HENRY THE FOURTH": true,
	"THE SECOND PART OF KING HENRY THE FOURTH": true,
	"THE LIFE OF KING HENRY THE FIFTH": true,
	"THE FIRST PART OF HENRY THE SIXTH": true,
	"THE SECOND PART OF KING HENRY THE SIXTH": true,
	"THE THIRD PART OF KING HENRY THE SIXTH": true,
	"KING HENRY THE EIGHTH": true,
	"KING JOHN": true,
	"THE TRAGEDY OF JULIUS CAESAR": true,
	"THE TRAGEDY OF KING LEAR": true,
	"LOVE’S LABOUR’S LOST": true,
	"THE TRAGEDY OF MACBETH": true,
	"MEASURE FOR MEASURE": true,
	"THE MERCHANT OF VENICE": true,
	"THE MERRY WIVES OF WINDSOR": true,
	"A MIDSUMMER NIGHT’S DREAM": true,
	"MUCH ADO ABOUT NOTHING": true,
	"THE TRAGEDY OF OTHELLO, MOOR OF VENICE": true,
	"PERICLES, PRINCE OF TYRE": true,
	"KING RICHARD THE SECOND": true,
	"KING RICHARD THE THIRD": true,
	"THE TRAGEDY OF ROMEO AND JULIET": true,
	"THE TAMING OF THE SHREW": true,
	"THE TEMPEST": true,
	"THE LIFE OF TIMON OF ATHENS": true,
	"THE TRAGEDY OF TITUS ANDRONICUS": true,
	"THE HISTORY OF TROILUS AND CRESSIDA": true,
	"TWELFTH NIGHT; OR, WHAT YOU WILL": true,
	"THE TWO GENTLEMEN OF VERONA": true,
	"THE TWO NOBLE KINSMEN": true,
	"THE WINTER’S TALE": true,
	"A LOVER’S COMPLAINT": true,
	"THE PASSIONATE PILGRIM": true,
	"THE PHOENIX AND THE TURTLE": true,
	"THE RAPE OF LUCRECE": true,
	"VENUS AND ADONIS": true,
}

type Shard struct {
	Title             string
	CompleteWorkShard string
	SuffixShard *suffixarray.Index
	CaseSuffixShard *suffixarray.Index
}

type FileLoadSharder struct {
	Shards []Shard
}

func NewFileLoader(filename string) (FileLoadSharder, error) {
	var f FileLoadSharder
	if file, err := os.Open(filename); err != nil {
		return FileLoadSharder{}, fmt.Errorf("open: %s", err)
	} else {
   		scanner := bufio.NewScanner(file)
   		s := Shard{
   			Title: "OTHER SECTION",
		}
   		for scanner.Scan() {
   			line := scanner.Text()
   			if bookTitles[line] {
   				s.SuffixShard = suffixarray.New([]byte(s.CompleteWorkShard))
				s.CaseSuffixShard = suffixarray.New([]byte(strings.ToLower(s.CompleteWorkShard)))
				f.Shards = append(f.Shards, s)
				s = Shard{Title: line}
			}
			s.CompleteWorkShard += line + "\n"
		}
		f.Shards = append(f.Shards, s)
		return f, nil
	}
}