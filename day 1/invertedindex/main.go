package main

import (
	"errors"
	"fmt"
	"strings"
)

type InvertedIndx struct {
	index map[string][]int
}

func NewInvertedIndx() *InvertedIndx {
	return &InvertedIndx{
		index: make(map[string][]int),
	}
}

func (ii *InvertedIndx) Add(docID int, txt string) error {
	tokenizedTxt, err := tokenise(&txt)
	if err != nil {
		return errors.New("invalid or empty string")
	}

	seen := make(map[string]bool)

	for _, word := range tokenizedTxt {
		if !seen[word] {
			ii.index[word] = append(ii.index[word], docID)
			seen[word] = true
		}
	}

	return nil
}

func (ii *InvertedIndx) Search(txt string) []int {
	return ii.index[strings.ToLower(txt)]
}

func tokenise(txt *string) ([]string, error) {
	if txt == nil || *txt == "" {
		return nil, errors.New("Invalid or Empty stirng")
	}

	newTxt := strings.ToLower(*txt)

	return strings.Fields(newTxt), nil
}

func main() {
	idx := NewInvertedIndx()

	idx.Add(1, "go is a fast programming language")
	idx.Add(2, "concurrence in go is powerful")
	idx.Add(3, "fast search engines use a inverted index")

	fmt.Println("search 'go': ", idx.Search("go"))
}
