package main

type WordDictionary struct {
	slice []string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {

}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

	for i:=0;i<len(this.slice);i++{


	}
}


func main() {
	t := Constructor()
	t.AddWord("bad")
	t.AddWord("dad")
	t.AddWord("mad")
	t.Search("pad") //false
	t.Search("bad") // true
	t.Search(".ad") // true
	t.Search("b..") // true

}
