package main

/*
https://leetcode.cn/problems/implement-trie-prefix-tree/solutions/
*/
type Trie struct {
    children [26]*Trie
    isEnd    bool
}

func Constructor() Trie {
    return Trie{}
}

func (t *Trie) Insert(word string) {

}

func (t *Trie) Search(word string) bool {

}

func (t *Trie) StartsWith(prefix string) bool {

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
