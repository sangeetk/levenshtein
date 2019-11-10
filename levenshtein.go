package levenshtein

// The Trie data structure keeps a set of words, organized with one node for
// each letter. Each node has a branch for each letter that may follow it in the
// set of words.
type TrieNode struct {
	Word     string
	Children map[string]*TrieNode
}

func New() *TrieNode {
	node := &TrieNode{}
	node.Children = make(map[string]*TrieNode)
	return node
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, letter := range []rune(word) {
		k := string(letter)
		if _, ok := t.Children[k]; !ok {
			node.Children[k] = New()
		}
		node = node.Children[k]
	}
	node.Word = word
}

// The search function returns a list of all words that are less than the given
// maximum distance from the target word
func (t *TrieNode) Search(word string, maxCost int) []string {
	// build first row
	currentRow := intRange(len(word) + 1)
	results := []string{}

	for k, v := range t.Children {
		v.searchr(k, word, currentRow, results, maxCost)
	}
	return results
}

// This recursive helper is used by the search function above. It assumes that
// the previousRow has been filled in already.
func (t *TrieNode) searchr(key, word string, previousRow []int, results []string, maxCost int) {
	columns := len(word) + 1
	currentRow := intRange(previousRow[0] + 1)

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
	if currentRow[len(currentRow)-1] <= maxCost && t.Word != "" {
		results = append(results, t.Word)
	}

	// if any entries in the row are less than the maximum cost, then
	// recursively search each branch of the trie
	if minCost(currentRow) <= maxCost {
		for k, v := range t.Children {
			v.searchr(k, word, currentRow, results, maxCost)
		}
	}
}

func intRange(size int) []int {
	var list []int
	for i := 0; i < size; i++ {
		list = append(list, i)
	}
	return list
}

func minCost(cost []int) int {
	min := 0
	for _, x := range cost {
		if x < min {
			min = x
		}
	}
	return min
}
