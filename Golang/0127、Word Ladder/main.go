func ladderLength(beginWord string, endWord string, wordList []string) int {
    //这个直接可以参考前一题的解法，直接一头一尾双向解析，连通即可
    wordDict := make(map[string]struct{}) //空的struct是不占空间的
    for _, v := range wordList{
        wordDict[v] = struct{}{}
    }
    if _,ok := wordDict[endWord]; !ok {
        return 0
    }
    
    isFound := false
    backend := false
    
    src := make(map[string]struct{})
    src[beginWord] = struct{}{}
    
    end := make(map[string]struct{})
    end[endWord] = struct{}{}
    
    paths := make(map[string][]string) //路径，k对应要解析的单词，v对应单词的由来路径，就是可以转换成k的合法单词
    
    nums := 0
    
    for len(src) > 0 && len(end) > 0 && !isFound {
        nums++
        
        if len(wordDict) == 0 {
            return 0
        }
        
        if len(src) > len(end) {
            //只遍历数目小的数组，提高效率
            src, end = end, src
            backend = !backend
        }
        
        newSrc := make(map[string]struct{})
        
        for k :=range src { 
            
            delete (wordDict, k)
            
            bytes := []byte(k) //准备逐位替换字符
            
            for i:=0; i < len(bytes); i++ {
                 for j:=0; j < 26; j++ {
                     if(isFound){
                         break
                     }
                     pre := bytes[i]
                     bytes[i] = byte('a' + j  )
                     newWord := string(bytes)
                     
                     if _,ok := end[newWord]; ok {
                        //打通了
                        isFound = true
                        target := k
                        solution := newWord

                         if backend {
                             target, solution = solution,target
                         }

                         if _,ok := paths[target]; ok {
                             paths[target] = append(paths[target], solution)
                         } else {
                             paths[target] = []string{solution}
                         }
                         
                         
                     } else {
                         fmt.Println(newWord)
                         //没找到，那就先看当前是不是合法单词
                         if _,ok := wordDict[newWord]; ok {
                             
                             newSrc[newWord] = struct{}{}
                             
                             //合法就存到paths总map中
                             target := k
                             solution := newWord
                             
                             if backend {
                                 target, solution = solution,target
                             }
                             
                             if _,ok := paths[target]; ok {
                                 paths[target] = append(paths[target], solution)
                             } else {
                                 paths[target] = []string{solution}
                             }
                         }
                     }
                     bytes[i] = pre
                }
                
            }
            
            
        }
        
        src = newSrc
    }
   
    if nums == 0 || !isFound {
        return 0
    }
    
    return nums+1
    
}




