func fullJustify(words []string, maxWidth int) []string {
    //一行一行的拼
    lens := len(words)
    res := make([]string, 0)
    
    nowNum :=0 //当前行累计的单词个数
    nowLens :=0 //当前单词的所有长度
    tmpRes := make([]string, 0)
    for i:=0; i < lens;{
        if len(words[i]) > maxWidth {
            break
        }
        if len(words[i]) + nowLens + nowNum <= maxWidth {
            nowNum++
            nowLens += len(words[i])
            tmpRes = append(tmpRes, words[i])
            i++
        } else {
            allWidth := maxWidth - nowLens //剩余可分配的空格宽度
            
            spaceNum := 0 //每个单词之间的空格长度
            reside := 0 //没分配完成的空格，要平均分配到左边的间隙中，每个间隙分配一个
            
            if nowNum > 1 {
                reside = allWidth % (nowNum-1)
                spaceNum =  allWidth/(nowNum-1)
            }
            rows := ""
            flag := 1
            
            rowLens := 0
            
            for j:=0; j < nowNum;{
                
                if flag == 1 {
                    rows  += tmpRes[j]
                    rowLens += len(tmpRes[j])
                    j++
                } else {
                    spaceLen := spaceNum
                    if reside > 0 {
                        spaceLen++
                        reside--
                    }
                    spaceStr := createSpace(spaceLen)
                    rowLens += spaceLen
                    rows += spaceStr
                }
                flag *= -1
            }
            
            if rowLens < maxWidth{
                rows += createSpace(maxWidth - rowLens)
            }
           
            res = append(res, rows)
            
            tmpRes = make([]string, 0)
            nowNum = 0
            nowLens = 0
        }
       
    }
    
     
    //处理最后一行
    if len(tmpRes) > 0{
       rows := ""
        flag := 1
        lens := 0
        for i:=0; i <len(tmpRes); {
            if flag == 1 {
                rows += tmpRes[i]
                lens += len(tmpRes[i])
                i++
            } else {
                rows += " "
                lens++
            }
            flag *= -1
        }
        rows += createSpace(maxWidth-lens)
        res = append(res, rows)
    }
    
    return res
}

func createSpace(num int) string {
    res := ""
    for i:=0; i < num;i++ {
        res += " "
    }
    return res
}
