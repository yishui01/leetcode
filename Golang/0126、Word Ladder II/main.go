func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    //Runtime: 20 ms, faster than 99.01%
    //我就想问问这解法怎么想出来的，这是人吗？这是人吗？这还是人吗？
    
    
    //思路就是迭代解析：
    //一头一尾两个单词存在两个map中，存在key里面，value是空的struct
    //迭代两个map中长度小的那个map，事实上这是优化手段，去掉也没事，只不过运行时间会从20ms变为140ms，去掉的话就是单向连通，不去掉就是双向连通
    //迭代的map叫src，要将src的单词转换成dst（另一个map）里的单词
    //src里的单词每次改变其中的一位为a-z的任意一个单词，看是否存在与dist中
    //不存在的话就在wordList中找是否是合法单词，合法的话就存到新的newSrc中，作为下一轮迭代，并将原本的target=>source映射存于paths这个map中，
    //newSrc迭代时成为新的src，此时将src与dist比较，如果src的长度大于dist，就将src和dist调换位置，然后进行迭代
    //就是要保证dist和src两边单词要连通
    
	wordDict := make(map[string]struct{})
	for _, word := range wordList {
		wordDict[word] = struct{}{}
	}
	// 提前返回
	if _, ok := wordDict[endWord]; !ok {
		return [][]string{}
	}

	src, dst := map[string]struct{}{}, map[string]struct{}{}
	src[beginWord] = struct{}{}
	dst[endWord] = struct{}{}

	found := false
	paths := make(map[string][]string)
	backward := false

	for len(src) != 0 && len(dst) != 0 && !found {

		// 始终遍历 小数组
		if len(src) > len(dst) {
			src, dst = dst, src
			backward = !backward
		}
		for w := range src {
			delete(wordDict, w)
		}

		newSrc := make(map[string]struct{})
		for word := range src {
			bytes := []byte(word) // 转成 []byte，然后修改值
			for i := 0; i < len(bytes); i++ {
				for j := 0; j < 26; j++ {
					bytes[i] = byte(j) + 'a' // 修改为 a-z
					source := word
					target := string(bytes) // 转成 string
					if _, ok := dst[target]; ok { // 已经连通
						if backward {
							source, target = target, source
						}
						if paths[target] == nil {
							paths[target] = make([]string, 0)
						}
						paths[target] = append(paths[target], source)
                         
                        fmt.Println("我是有的---",i,"---",source,"---",target,"---",backward)
                        
						found = true
					} else {
						if _, ok := wordDict[target]; ok {
							newSrc[target] = struct{}{}
							if backward {
								source, target = target, source
							}
							if paths[target] == nil {
								paths[target] = make([]string, 0)
							}
                            
                           fmt.Println(i,"---",source,"---",target,"---",backward)
                            
							paths[target] = append(paths[target], source)
						}
					}
				}
				// 恢复原来的值
				bytes[i] = word[i]
			}
		}
		src = newSrc
	}
    
	ans := make([][]string, 0)
    
    
    
	getPath(paths, beginWord, endWord, []string{endWord}, &ans)

	return ans
}
func getPath(parents map[string][]string, beginWord, cur string, path []string, ans *[][]string) {
	if cur == beginWord {
		newPath := make([]string, len(path))
		copy(newPath, path)
		for i, j := 0, len(newPath)-1; i < j; i, j = i+1, j-1 {
			newPath[i], newPath[j] = newPath[j], newPath[i]
		}
		*ans = append(*ans, newPath)
		return
	}

   // fmt.Println("*********",parents[cur])
	for _, p := range parents[cur] {
		path = append(path, p)
		getPath(parents, beginWord, p, path, ans)
		path = path[:len(path)-1]
	}
}
