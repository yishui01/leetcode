func longestCommonPrefix(strs []string) string {
    //生成一个公共前缀数组，一直比对,直到循环比对结束或者数组减为空
    comm := make(map[int]byte)
    lens := len(strs)
    i:=0
    if lens == 0 {
        return ""
    }
    if lens == 1 {
        return strs[0]
    }
    for {
        if i >= len(strs[0]) || i >= len(strs[1]) || strs[0][i] != strs[1][i] {
            break
        }
        comm[i] = strs[0][i]
        i++
    }
    for i:=2;i<lens;i++{
        if len(strs[i]) == 0 {
            return ""
        }
        for j:=0;j<len(strs[i]);j++{
            if j>= len(comm) {
                break
            } else {
                val,ok := comm[j]
                if ok && val == strs[i][j] {
                    if j == len(strs[i]) - 1 {
                        comm[j+1] = '0'
                    }
                    continue
                } else {
                     if j == 0 {
                        return ""
                    }
                    comm[j]='0'
                    break
                }
            }
        }
    }
   
    
    strArr := make([]byte,0)
    
    for i:=0;i<len(comm);i++ {
        val,_ := comm[i]
        if val == '0' {
            return string(strArr)
        }
        strArr = append(strArr, val)
    }

    return string(strArr)
}