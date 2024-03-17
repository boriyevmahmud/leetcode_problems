package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Mahmud"
	prefix := "Mah"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search(word)
	fmt.Println("param_2: ", param_2)
	param_3 := obj.StartsWith(prefix)
	fmt.Println("param_3: ", param_3)

}

type Trie struct {
	Words            []string
	LastInsertedWord string
}

func Constructor() Trie {
	return Trie{
		Words:            []string{},
		LastInsertedWord: "",
	}

}

func (this *Trie) Insert(word string) {
	this.Words = append(this.Words, word)
	this.LastInsertedWord = word

}

func (this *Trie) Search(word string) bool {
	if this.LastInsertedWord == word {
		return true
	}
	for _, w := range this.Words {
		if w == word {
			return true
		}
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return strings.HasPrefix(this.LastInsertedWord, prefix)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
