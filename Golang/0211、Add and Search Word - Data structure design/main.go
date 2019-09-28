//就用之前的前缀树的方法就行
type WordDictionary struct {
    Next [26]*WordDictionary
    Word string
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{Next:[26]*WordDictionary{}}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    node := this
    for i:=0;i<len(word);i++{
        nowChar := int(word[i] - 'a')
        node.Next[nowChar] = &WordDictionary{Next:[26]*WordDictionary{}}
        node = node.Next[nowChar]
    }
    node.Word = word
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    node := this
    for i:=0; i < len(word); i++ {
        if word[i] == '*' {
            continue
        }
        if word[i] == '.'{
            for _,v :=range node.Next {
                if v != nil {
                    b := []byte(word)
                    for j:=0;j<=i;j++{
                         b[j] = '*'
                    }
                    if v.Search(string(b)){
                       return true
                 }
              }
            }
            return false
        } else {
           nowChar := int(word[i]-'a')
            if node.Next[nowChar] == nil {
                return false
            }
            node = node.Next[nowChar]
        }
    }
    
    fmt.Println(word,node.Word)
    if len(word) == len(node.Word) {
        for i:=0;i < len(word);i++{
            if word[i] == '*' {
                continue
            }
            if word[i] != node.Word[i] {
                return false
            }
        }
        return true
    }
    
    return false
    
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
