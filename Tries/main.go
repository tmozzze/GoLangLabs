package main

import "fmt"

type Node struct {
	Children map[rune]*Node
	isEnd    bool
	Value    int
}

type Trie struct {
	Root *Node
}

func NewNode() *Node {
	return &Node{Children: make(map[rune]*Node)}
}

func NewTrie() *Trie {
	return &Trie{Root: &Node{Children: make(map[rune]*Node)}}
}

func (trie *Trie) Insert(word string, value int) {
	current := trie.Root
	for _, char := range word {
		if current.Children[char] == nil {
			current.Children[char] = NewNode()
		}
		current = current.Children[char]
	}
	current.isEnd = true
	current.Value = value
}

func (trie *Trie) Search(word string) int {
	current := trie.Root
	for _, char := range word {
		if current.Children[char] == nil {
			return 999999
		}
		current = current.Children[char]
	}
	if current.isEnd {
		return current.Value
	}
	return 999999
}

func (trie *Trie) Delete(word string) {
	current := trie.Root
	for _, char := range word {
		if current.Children[char] == nil {
			fmt.Println("Trie Delete Not Found")
			return
		}
		current = current.Children[char]
	}
	if current.isEnd {
		current.isEnd = false
		fmt.Println("Trie Delete Found")
		return
	}
	fmt.Println("Trie Delete not Found")
	return
}

func main() {
	trie := NewTrie()
	trie.Insert("joj", 2)
	fmt.Println(trie.Search("joj"))
	trie.Delete("joj")
	fmt.Println(trie.Search("joj"))
}
