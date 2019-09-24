type Trie struct {
    Next [26]*Trie
    Word string
}

//前缀树，那每一层都要有对应的字母索引才行，一层一层往下通

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{Next:[26]*Trie{}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this
    for i:=0;i<len(word);i++{
        nowChar := int(word[i]-'a')
        if node.Next[nowChar] == nil {
            node.Next[nowChar] = &Trie{Next:[26]*Trie{}}
        }
        if i == len(word)-1 {
            node.Next[nowChar].Word = word
        }
        node = node.Next[nowChar]
    }
  
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this
    for i:=0; i < len(word);i++{
        nowChar := int(word[i]-'a')
        if node.Next[nowChar] == nil {
            return false
        }
        node = node.Next[nowChar]
    }
    
    if node.Word == word {
        return true
    }
    
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this
    for i:=0; i < len(prefix);i++{
        nowChar := int(prefix[i]-'a')
        if node.Next[nowChar] == nil {
            return false
        }
        node = node.Next[nowChar]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
