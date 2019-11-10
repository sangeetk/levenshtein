// Levenshtein Distance in Golang
package levenshtein

import (
	"fmt"
	"testing"
)

const DICTIONARY = "/usr/share/dict/words"

func TestLevenstein(t *testing.T) {

	lev := New()

	lev.Insert("sangeet")
	results := lev.Search("sanjeet", 1)

	for _, match := range results {
		fmt.Println("match:", match)
	}
}
