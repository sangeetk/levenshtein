// Levenshtein Distance in Golang
package main

import (
	"fmt"

	"github.com/sangeetk/levenshtein"
)

const DICTIONARY = "/usr/share/dict/words"

func main() {

	lev := levenshtein.New()

	lev.Insert("sangeet")
	results := lev.Search("sanjeet", 1)

	for _, match := range results {
		fmt.Println("match:", match)
	}
}
