package levenshtein

// The Trie data structure keeps a set of words, organized with one node for
// each letter. Each node has a branch for each letter that may follow it in the
// set of words.
type Levenshtein struct {
	Word     string
	Children map[string]*Levenshtein
}

type Result struct {
	Word string
	Cost int
}

func New() *Levenshtein {
	node := &Levenshtein{}
	node.Children = make(map[string]*Levenshtein)
	return node
}

func (l *Levenshtein) Insert(word string) {
	node := l
	for _, letter := range []rune(word) {
		k := string(letter)
		if _, ok := node.Children[k]; !ok {
			node.Children[k] = New()
		}
		node = node.Children[k]
	}
	node.Word = word
}

// The search function returns a list of all words that are less than the given
// maximum distance from the target word
func (l *Levenshtein) Search(word string, maxCost int) []Result {
	// build first row
	currentRow := initRange(len(word) + 1)
	results := []Result{}

	for k, v := range l.Children {
		r := v.searchr(k, word, currentRow, maxCost)
		for _, result := range r {
			results = append(results, result)
		}
	}
	return results
}

// This recursive helper is used by the search function above. It assumes that
// the previousRow has been filled in already.
func (l *Levenshtein) searchr(key, word string, previousRow []int, maxCost int) []Result {
	columns := len(word) + 1
	currentRow := initRange(previousRow[0] + 1)
	results := []Result{}

	// Build one row for the letter, with a column for each letter in the target
	// word, plus one for the empty string at column 0
	for column := 1; column < columns; column++ {
		insertCost := currentRow[column-1] + 1
		deleteCost := previousRow[column] + 1
		replaceCost := 0

		letters := []rune(word)
		if string(letters[column-1]) != key {
			replaceCost = previousRow[column-1] + 1
		} else {
			replaceCost = previousRow[column-1]
		}

		currentRow = append(currentRow, minCost([]int{insertCost, deleteCost, replaceCost}))
	}

	// if the last entry in the row indicates the optimal cost is less than the
	// maximum cost, and there is a word in this trie node, then add it.
	if currentRow[len(currentRow)-1] <= maxCost && l.Word != "" {
		r := Result{Word: l.Word, Cost: currentRow[len(currentRow)-1]}
		results = append(results, r)
	}

	// if any entries in the row are less than the maximum cost, then
	// recursively search each branch of the trie
	if minCost(currentRow) <= maxCost {
		for k, v := range l.Children {
			r := v.searchr(k, word, currentRow, maxCost)
			for _, result := range r {
				results = append(results, result)
			}
		}
	}
	return results
}

func initRange(size int) []int {
	var list []int
	for i := 0; i < size; i++ {
		list = append(list, i)
	}
	return list
}

func minCost(cost []int) int {
	var min int
	for i, x := range cost {
		if i == 0 || x < min {
			min = x
		}
	}
	return min
}
