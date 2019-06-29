func minWindow(s string, t string) string {
    //滑动窗口 //leftTopIndex  和  leftSecondIndex 
    //右边界一直扩大, 遇到和左边界重复的字母，就收缩左边界到下一个leftSecondIndex,这里leftSecondIndex要注意，在设置的时候，如果遇到连续个leftSecondIndex，要更新为最后一个字母的index
    //这样收缩左边框才能缩小到最小
    //右边框扩大规则为非字母串的字母直接扩大，进入下一轮，遇到非左边界的字母，看看当前是否在second循环模式，在的话又有leftScond，就leftSecondIndex++，不然就扩大，记录res
    //如果res到了规定的长度，就是结果，但是别急，要遍历到最后
    
    leftTopIndex := -1
    var leftLetter byte
   
    rightIndex := 0
    
    lens := len(s)
    subLens := len(t)
    
    resNum :=0
    
    min := lens+1
    resLeft := -1
    resRight := -1
    
    
    if lens == 0 || subLens == 0 || subLens > lens {
        return ""
    }
    
    maps := make(map[byte]int) //遇到一个弹出一个,弹完了就是找到结果了
    constMap := make(map[byte]int)
    // indexSlice := make([]int, 0) //记录每个OK的字符出现的index
    
    for _,v := range t {
        maps[byte(v)]++
        constMap[byte(v)] = 0
    }
    
    fmt.Println(maps)
    
  
    for i:=0;i < lens; i++ {
        
        if _,ok := constMap[s[i]];ok {
            
            maps[s[i]]--
            constMap[s[i]]++
            
            if val,ok := maps[s[i]]; ok && val >= 0 {
                resNum++
            }
            if leftTopIndex == -1 {
                leftTopIndex = i
            }
            
            if leftLetter == 0 {
                leftLetter = s[i]
            } else if s[i] == leftLetter { //看和最左边的是不是相等
                if _,ok := maps[s[i]]; ok {
                    //判断是否要收缩左边框
                    if maps[s[i]] <= 0 {
                        //收缩,要收缩到下一个有效字母
                        leftTopIndex++
                        for leftTopIndex < lens {
                                if _,ok := maps[s[leftTopIndex]]; ok {
                                    if val,ok2 := constMap[s[leftTopIndex]];ok2 && val > 0 {
                                        constMap[s[leftTopIndex]]--
                                        leftTopIndex++
                                        continue
                                    } else {
                                        break
                                    }
                                }
                              leftTopIndex++
                        }
                         leftLetter = s[leftTopIndex-1]
                    }

                }
                
            }
            
        }
        
        rightIndex++
        
         fmt.Println(leftTopIndex,"---",rightIndex,"---",min, "---",len(maps),"---",resNum)
        
        if resNum == subLens && rightIndex - leftTopIndex < min {
            min = rightIndex - leftTopIndex
            resLeft = leftTopIndex
            resRight = rightIndex
        }
        
    }
    fmt.Println(rightIndex,"---",leftTopIndex)
    if leftTopIndex == -1 || resLeft == -1 {
        return "666"
    }
    
    return s[resLeft:resRight]
    
}
