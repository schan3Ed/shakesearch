package model

import (
	"pulley.com/shakesearch/helper"
	"testing"
)

func TestFindPrevious(t *testing.T) {
	article := `
	Space\n
	Space\n
	Space\n
	Space\n
	You are reading my test\n
	That's nice\n
	I hope you enjoy reading my code\n
	LET ME IN\n
	`
	prev := findPrevious(article, 100)
	helper.AssertEqual(t, prev, 28)
	prev = findPrevious(article, 110)
	helper.AssertEqual(t, prev, 28)
	prev = findPrevious(article, 30)
	helper.AssertEqual(t, prev, 0)
}


func TestFindAfter(t *testing.T) {
	article := `
	Space\n
	Space\n
	Space\n
	Space\n
	You are reading my test\n
	That's nice\n
	I hope you enjoy reading my code\n
	LET ME IN\n
	`
	after := findAfter(article, 100)
	helper.AssertEqual(t, after, 128)
	after = findAfter(article, 110)
	helper.AssertEqual(t, after, 128)
	after = findAfter(article, 30)
	helper.AssertEqual(t, after, 114)
	after = findAfter(article, 1)
	helper.AssertEqual(t, after, 36)
	after = findAfter(article, 3)
	helper.AssertEqual(t, after, 36)
}

