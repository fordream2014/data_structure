package algrithm

import "fmt"

type TrieNode struct {
	links []*TrieNode
	isEnd bool
	size int
}

func (this *TrieNode) containKey(char byte) bool {
	ind := int(char - 'a')
	return this.links[ind] != nil
}

func (this *TrieNode) getNode(char byte) *TrieNode {
	ind := int(char - 'a')
	return this.links[ind]
}

func (this *TrieNode) insert(char byte) {
	ind := int(char - 'a')
	this.links[ind] = &TrieNode{
		links: make([]*TrieNode, 26),
	}
	this.size ++
}

func (this *TrieNode) getLinks() int {
	return this.size
}

func (this *TrieNode) end() bool {
	return this.isEnd
}

type Trie struct {
	root *TrieNode
}

func generateTrie() *Trie{
	root := &TrieNode{
		links: make([]*TrieNode, 26),
	}
	return &Trie{
		root: root,
	}
}

func (this *Trie) insert(str string) bool {
	if len(str) == 0 {
		return false
	}
	curNode := this.root
	for i:=0; i<len(str); i++ {
		if curNode.containKey(str[i]) {
			curNode = curNode.getNode(str[i])
		} else {
			curNode.insert(str[i])
			curNode = curNode.getNode(str[i])

			if i == len(str) - 1 {
				curNode.isEnd = true
			}
		}
	}
	return true
}

func (this *Trie) findLongestPrefix(word string) string {
	cur := this.root
	result := ""
	for i:=0; i<len(word); i++ {
		if cur.containKey(word[i]) && cur.getLinks() == 1 && !cur.end() {
			result += string(word[i])
			cur = cur.getNode(word[i])
		} else {
			break
		}
	}
	return result
}

func TestTrie() {
	str0 := "abc"
	str1 := "abcas"
	str2 := "abcsdfw"
	trie := generateTrie()
	//trie.insert(str0)
	trie.insert(str1)
	trie.insert(str2)

	prefix := trie.findLongestPrefix(str0)
	fmt.Println(prefix)
}
