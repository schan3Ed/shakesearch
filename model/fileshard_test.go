package model

import (
	"pulley.com/shakesearch/helper"
	"testing"
)

func TestLoadFile(t *testing.T) {
	f, err := NewFileLoader("../load_test.txt")
	helper.AssertEqual(t, err, nil)
	helper.AssertEqual(t, len(f.Shards), 3)
	helper.AssertEqual(t, f.Shards[0].Title, "OTHER SECTION")
	helper.AssertEqual(t, f.Shards[1].Title, "THE TEMPEST")
	helper.AssertEqual(t, f.Shards[2].Title, "KING JOHN")
}

