package levenshtein

import (
	"fmt"
	"testing"
)

func TestLevenstein(t *testing.T) {

	lev := New()
	lev.Insert("levenshtein")
	results := lev.Search("lavenstein", 2)

	for _, match := range results {
		fmt.Println("match:", match)
	}
}
