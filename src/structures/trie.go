package structures

import (
	"fmt"
	"strings"
)

type TNode struct {
	value rune
	next  []TNode
	isEnd bool
}

type Trie struct {
	head  TNode
	words []string
}

func NewTrie(words []string) *Trie {
	head := TNode{value: ' ', next: make([]TNode, 0), isEnd: false}
	trie := Trie{words: words, head: head}
	trie.build()
	return &trie
}

// Build
func (self *Trie) build() {
	for idx := range self.words {
		self.addWord(&self.words[idx])
	}
}

func (self *Trie) addWord(word *string) {
	currentNode := &self.head
	for _, char := range *word {
		next := currentNode.getNext(char)
		if next == nil {
			currentNode.next = append(currentNode.next, TNode{value: char, next: make([]TNode, 0), isEnd: false})
			currentNode = &currentNode.next[len(currentNode.next)-1]
		} else {
			currentNode = next
		}
	}
	currentNode.isEnd = true
}

func (self *TNode) getNext(char rune) *TNode {
	for nIdx := range self.next {
		if self.next[nIdx].value == char {
			return &self.next[nIdx]
		}
	}
	return nil
}

// Matches
func (self *Trie) GetMatches(str string) []string {
	matches := make([]string, 0)

	for startIdx := range str {
		endIdx := startIdx
		currentNode := self.head
		for {
			if currentNode.isEnd {
				matches = append(matches, str[startIdx:endIdx])
				break
			}
			if endIdx == len(str) {
				break
			}
			var next *TNode
			for idx := range currentNode.next {
				if currentNode.next[idx].value == rune(str[endIdx]) {
					next = &currentNode.next[idx]
					break
				}
			}
			if next != nil {
				endIdx += 1
				currentNode = *next
			} else {
				break
			}
		}
	}

	return matches
}

// Debug
func (self *Trie) print() {
	for i := range self.head.next {
		self.head.next[i].print(0)
	}
}

func (self *TNode) print(depth int) {
	println(fmt.Sprintf("%s%s  %v", strings.Repeat("  ", depth), string(self.value), self.isEnd))
	for i := range self.next {
		self.next[i].print(depth + 1)
	}
}
