package searchengine

import (
	"strings"
)

type TextEngine struct {
	index map[string][]int
	docs  []string
}

func NewTextEngine() *TextEngine {
	return &TextEngine{
		index: make(map[string][]int),
		docs:  []string{},
	}
}
func (te *TextEngine) AddDocument(doc string) {
	te.docs = append(te.docs, doc)
	docID := len(te.docs) - 1
	te.TokenizeAndIndex(doc, docID)

}
func (te *TextEngine) TokenizeAndIndex(doc string, docID int) {
	words := strings.Fields(doc)
	for _, word := range words {
		word := strings.ToLower(word)
		_, ok := te.index[word]
		if !ok {
			te.index[word] = []int{docID}
		} else {
			te.index[word] = append(te.index[word], docID)
		}

	}
}
func Search() {

}
